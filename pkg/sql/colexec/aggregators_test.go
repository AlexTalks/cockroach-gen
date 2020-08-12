// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/cockroachdb/apd/v2"
	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/col/coldatatestutils"
	"github.com/cockroachdb/cockroach/pkg/settings/cluster"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase"
	"github.com/cockroachdb/cockroach/pkg/sql/colmem"
	"github.com/cockroachdb/cockroach/pkg/sql/execinfrapb"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/cockroach/pkg/testutils/skip"
	"github.com/cockroachdb/cockroach/pkg/util/duration"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/log"
	"github.com/cockroachdb/cockroach/pkg/util/randutil"
	"github.com/stretchr/testify/require"
)

var (
	defaultGroupCols = []uint32{0}
	defaultAggCols   = [][]uint32{{1}}
	defaultAggFns    = []execinfrapb.AggregatorSpec_Func{execinfrapb.AggregatorSpec_SUM_INT}
	defaultTyps      = []*types.T{types.Int, types.Int}
)

type aggregatorTestCase struct {
	// typs, aggFns, groupCols, and aggCols will be set to their default
	// values before running a test if nil.
	typs           []*types.T
	aggFns         []execinfrapb.AggregatorSpec_Func
	groupCols      []uint32
	aggCols        [][]uint32
	constArguments [][]execinfrapb.Expression
	// spec will be populated during init().
	spec           *execinfrapb.AggregatorSpec
	input          tuples
	unorderedInput bool
	expected       tuples
	// {output}BatchSize() if not 0 are passed in to NewOrderedAggregator to
	// divide input/output batches.
	batchSize       int
	outputBatchSize int
	name            string

	// convToDecimal will convert any float64s to apd.Decimals. If a string is
	// encountered, a best effort is made to convert that string to an
	// apd.Decimal.
	convToDecimal bool
}

// aggType is a helper struct that allows tests to test both the ordered and
// hash aggregators at the same time.
type aggType struct {
	new func(
		allocator *colmem.Allocator,
		input colexecbase.Operator,
		inputTypes []*types.T,
		spec *execinfrapb.AggregatorSpec,
		evalCtx *tree.EvalContext,
		constructors []execinfrapb.AggregateConstructor,
		constArguments []tree.Datums,
		outputTypes []*types.T,
		isScalar bool,
	) (colexecbase.Operator, error)
	name string
}

var aggTypes = []aggType{
	{
		// This is a wrapper around NewHashAggregator so its signature is compatible
		// with orderedAggregator.
		new: func(
			allocator *colmem.Allocator,
			input colexecbase.Operator,
			inputTypes []*types.T,
			spec *execinfrapb.AggregatorSpec,
			evalCtx *tree.EvalContext,
			constructors []execinfrapb.AggregateConstructor,
			constArguments []tree.Datums,
			outputTypes []*types.T,
			_ bool,
		) (colexecbase.Operator, error) {
			return NewHashAggregator(allocator, input, inputTypes, spec, evalCtx, constructors, constArguments, outputTypes)
		},
		name: "hash",
	},
	{
		new:  NewOrderedAggregator,
		name: "ordered",
	},
}

func (tc *aggregatorTestCase) init() error {
	if tc.convToDecimal {
		for _, tuples := range []tuples{tc.input, tc.expected} {
			for _, tuple := range tuples {
				for i, e := range tuple {
					switch v := e.(type) {
					case float64:
						d := &apd.Decimal{}
						d, err := d.SetFloat64(v)
						if err != nil {
							return err
						}
						tuple[i] = *d
					case string:
						d := &apd.Decimal{}
						d, _, err := d.SetString(v)
						if err != nil {
							// If there was an error converting the string to decimal, just
							// leave the datum as is.
							continue
						}
						tuple[i] = *d
					}
				}
			}
		}
	}
	if tc.groupCols == nil {
		tc.groupCols = defaultGroupCols
	}
	if tc.aggFns == nil {
		tc.aggFns = defaultAggFns
	}
	if tc.aggCols == nil {
		tc.aggCols = defaultAggCols
	}
	if tc.typs == nil {
		tc.typs = defaultTyps
	}
	if tc.batchSize == 0 {
		tc.batchSize = coldata.BatchSize()
	}
	if tc.outputBatchSize == 0 {
		tc.outputBatchSize = coldata.BatchSize()
	}
	aggregations := make([]execinfrapb.AggregatorSpec_Aggregation, len(tc.aggFns))
	for i, aggFn := range tc.aggFns {
		aggregations[i].Func = aggFn
		aggregations[i].ColIdx = tc.aggCols[i]
		if tc.constArguments != nil {
			aggregations[i].Arguments = tc.constArguments[i]
		}
	}
	tc.spec = &execinfrapb.AggregatorSpec{
		GroupCols:    tc.groupCols,
		Aggregations: aggregations,
	}
	return nil
}

func TestAggregatorOneFunc(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	testCases := []aggregatorTestCase{
		{
			input: tuples{
				{0, 1},
			},
			expected: tuples{
				{1},
			},
			name:            "OneTuple",
			outputBatchSize: 4,
		},
		{
			input: tuples{
				{0, 1},
				{0, 1},
			},
			expected: tuples{
				{2},
			},
			name: "OneGroup",
		},
		{
			input: tuples{
				{0, 1},
				{0, 0},
				{0, 1},
				{1, 4},
				{2, 5},
			},
			expected: tuples{
				{2},
				{4},
				{5},
			},
			batchSize: 2,
			name:      "MultiGroup",
		},
		{
			input: tuples{
				{0, 1},
				{0, 2},
				{0, 3},
				{1, 4},
				{1, 5},
			},
			expected: tuples{
				{6},
				{9},
			},
			batchSize: 1,
			name:      "CarryBetweenInputBatches",
		},
		{
			input: tuples{
				{0, 1},
				{0, 2},
				{0, 3},
				{0, 4},
				{1, 5},
				{2, 6},
			},
			expected: tuples{
				{10},
				{5},
				{6},
			},
			batchSize:       2,
			outputBatchSize: 1,
			name:            "CarryBetweenOutputBatches",
		},
		{
			input: tuples{
				{0, 1},
				{0, 1},
				{1, 2},
				{2, 3},
				{2, 3},
				{3, 4},
				{3, 4},
				{4, 5},
				{5, 6},
				{6, 7},
				{7, 8},
			},
			expected: tuples{
				{2},
				{2},
				{6},
				{8},
				{5},
				{6},
				{7},
				{8},
			},
			batchSize:       3,
			outputBatchSize: 1,
			name:            "CarryBetweenInputAndOutputBatches",
		},
		{
			input: tuples{
				{0, 1},
				{0, 2},
				{0, 3},
				{0, 4},
			},
			expected: tuples{
				{10},
			},
			batchSize:       1,
			outputBatchSize: 1,
			name:            "NoGroupingCols",
			groupCols:       []uint32{},
		},
		{
			input: tuples{
				{1, 0, 0},
				{2, 0, 0},
				{3, 0, 0},
				{4, 0, 0},
			},
			expected: tuples{
				{10},
			},
			batchSize:       1,
			outputBatchSize: 1,
			name:            "UnusedInputColumns",
			typs:            []*types.T{types.Int, types.Int, types.Int},
			groupCols:       []uint32{1, 2},
			aggCols:         [][]uint32{{0}},
		},
		{
			input: tuples{
				{nil, 1},
				{4, 42},
				{nil, 2},
			},
			expected: tuples{
				{3},
				{42},
			},
			name:           "UnorderedWithNullsInGroupingCol",
			unorderedInput: true,
		},
		{
			aggFns: []execinfrapb.AggregatorSpec_Func{
				execinfrapb.AggregatorSpec_ANY_NOT_NULL,
				execinfrapb.AggregatorSpec_COUNT_ROWS,
			},
			aggCols:        [][]uint32{{0}, {}},
			typs:           []*types.T{types.Int},
			unorderedInput: true,
			input: tuples{
				{1},
				{2},
				{1},
				{nil},
				{3},
				{1},
				{3},
				{4},
				{1},
				{nil},
				{2},
				{4},
				{2},
			},
			expected: tuples{
				{nil, 2},
				{1, 4},
				{2, 3},
				{3, 2},
				{4, 2},
			},
		},
	}

	evalCtx := tree.MakeTestingEvalContext(cluster.MakeTestingClusterSettings())
	defer evalCtx.Stop(context.Background())
	// Run tests with deliberate batch sizes and no selection vectors.
	ctx := context.Background()
	for _, tc := range testCases {
		if err := tc.init(); err != nil {
			t.Fatal(err)
		}

		constructors, constArguments, outputTypes, err := ProcessAggregations(
			&evalCtx, nil /* semaCtx */, tc.spec.Aggregations, tc.typs,
		)
		require.NoError(t, err)
		if !tc.unorderedInput {
			log.Infof(ctx, "%s", tc.name)
			tupleSource := newOpTestInput(tc.batchSize, tc.input, tc.typs)
			a, err := NewOrderedAggregator(
				testAllocator, tupleSource, tc.typs, tc.spec, &evalCtx,
				constructors, constArguments, outputTypes, false, /* isScalar */
			)
			if err != nil {
				t.Fatal(err)
			}

			out := newOpTestOutput(a, tc.expected)
			// Explicitly reinitialize the aggregator with the given output batch
			// size.
			a.(*orderedAggregator).initWithInputAndOutputBatchSize(tc.batchSize, tc.outputBatchSize)
			if err := out.VerifyAnyOrder(); err != nil {
				t.Fatal(err)
			}
		}

		// Run randomized tests on this test case.
		for _, agg := range aggTypes {
			if tc.unorderedInput && agg.name == "ordered" {
				// This test case has unordered input, so we skip ordered
				// aggregator.
				continue
			}
			log.Infof(ctx, "%s/Randomized/%s", tc.name, agg.name)
			runTestsWithTyps(t, []tuples{tc.input}, [][]*types.T{tc.typs}, tc.expected, unorderedVerifier,
				func(input []colexecbase.Operator) (colexecbase.Operator, error) {
					return agg.new(
						testAllocator, input[0], tc.typs, tc.spec, &evalCtx,
						constructors, constArguments, outputTypes, false, /* isScalar */
					)
				})
		}
	}
}

func TestAggregatorMultiFunc(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	// TODO(yuzefovich): introduce nicer aliases for the protobuf generated
	// ones and use those throughout the codebase.
	avgFn := execinfrapb.AggregatorSpec_AVG
	testCases := []aggregatorTestCase{
		{
			aggFns: []execinfrapb.AggregatorSpec_Func{execinfrapb.AggregatorSpec_SUM_INT, execinfrapb.AggregatorSpec_SUM_INT},
			aggCols: [][]uint32{
				{2}, {1},
			},
			input: tuples{
				{0, 1, 2},
				{0, 1, 2},
			},
			typs: []*types.T{types.Int, types.Int, types.Int},
			expected: tuples{
				{4, 2},
			},
			name: "OutputOrder",
		},
		{
			aggFns: []execinfrapb.AggregatorSpec_Func{execinfrapb.AggregatorSpec_SUM, execinfrapb.AggregatorSpec_SUM_INT},
			aggCols: [][]uint32{
				{2}, {1},
			},
			input: tuples{
				{0, 1, 1.3},
				{0, 1, 1.6},
				{0, 1, 0.5},
				{1, 1, 1.2},
			},
			typs: []*types.T{types.Int, types.Int, types.Decimal},
			expected: tuples{
				{3.4, 3},
				{1.2, 1},
			},
			name:          "SumMultiType",
			convToDecimal: true,
		},
		{
			aggFns: []execinfrapb.AggregatorSpec_Func{execinfrapb.AggregatorSpec_AVG, execinfrapb.AggregatorSpec_SUM},
			aggCols: [][]uint32{
				{1}, {1},
			},
			input: tuples{
				{0, 1.1},
				{0, 1.2},
				{0, 2.3},
				{1, 6.21},
				{1, 2.43},
			},
			typs: []*types.T{types.Int, types.Decimal},
			expected: tuples{
				{"1.5333333333333333333", 4.6},
				{4.32, 8.64},
			},
			name:          "AvgSumSingleInputBatch",
			convToDecimal: true,
		},
		{
			aggFns: []execinfrapb.AggregatorSpec_Func{
				execinfrapb.AggregatorSpec_BOOL_AND,
				execinfrapb.AggregatorSpec_BOOL_OR,
			},
			aggCols: [][]uint32{
				{1}, {1},
			},
			input: tuples{
				{0, true},
				{1, false},
				{2, true},
				{2, false},
				{3, true},
				{3, true},
				{4, false},
				{4, false},
				{5, false},
				{5, nil},
				{6, nil},
				{6, true},
				{7, nil},
				{7, false},
				{7, true},
				{8, nil},
				{8, nil},
			},
			typs: []*types.T{types.Int, types.Bool},
			expected: tuples{
				{true, true},
				{false, false},
				{false, true},
				{true, true},
				{false, false},
				{false, false},
				{true, true},
				{false, true},
				{nil, nil},
			},
			name: "BoolAndOrBatch",
		},
		{
			aggFns: []execinfrapb.AggregatorSpec_Func{
				execinfrapb.AggregatorSpec_ANY_NOT_NULL,
				execinfrapb.AggregatorSpec_ANY_NOT_NULL,
				execinfrapb.AggregatorSpec_ANY_NOT_NULL,
				execinfrapb.AggregatorSpec_MIN,
				execinfrapb.AggregatorSpec_SUM,
			},
			input: tuples{
				{2, 1.0, "1.0", 2.0},
				{2, 1.0, "1.0", 4.0},
				{2, 2.0, "2.0", 6.0},
			},
			expected: tuples{
				{2, 1.0, "1.0", 2.0, 6.0},
				{2, 2.0, "2.0", 6.0, 6.0},
			},
			batchSize: 1,
			typs:      []*types.T{types.Int, types.Decimal, types.Bytes, types.Decimal},
			name:      "MultiGroupColsWithPointerTypes",
			groupCols: []uint32{0, 1, 2},
			aggCols: [][]uint32{
				{0}, {1}, {2}, {3}, {3},
			},
		},
		{
			aggFns: []execinfrapb.AggregatorSpec_Func{
				execinfrapb.AggregatorSpec_ANY_NOT_NULL,
				execinfrapb.AggregatorSpec_SUM_INT,
			},
			input: tuples{
				{`{"id": null}`, -1},
				{`{"id": 0, "data": "s1"}`, 1},
				{`{"id": 0, "data": "s1"}`, 2},
				{`{"id": 1, "data": "s2"}`, 10},
				{`{"id": 1, "data": "s2"}`, 11},
				{`{"id": 2, "data": "s3"}`, 100},
				{`{"id": 2, "data": "s3"}`, 101},
				{`{"id": 2, "data": "s4"}`, 102},
			},
			expected: tuples{
				{`{"id": null}`, -1},
				{`{"id": 0, "data": "s1"}`, 3},
				{`{"id": 1, "data": "s2"}`, 21},
				{`{"id": 2, "data": "s3"}`, 201},
				{`{"id": 2, "data": "s4"}`, 102},
			},
			typs:      []*types.T{types.Jsonb, types.Int},
			name:      "GroupOnJsonColumns",
			groupCols: []uint32{0},
			aggCols: [][]uint32{
				{0}, {1},
			},
		},
		{
			input: tuples{
				{0, nil, 1, 1, 1.0, 1.0, duration.MakeDuration(1, 1, 1)},
				{0, 1, nil, 2, 2.0, 2.0, duration.MakeDuration(2, 2, 2)},
				{0, 2, 2, nil, 3.0, 3.0, duration.MakeDuration(3, 3, 3)},
				{0, 3, 3, 3, nil, 4.0, duration.MakeDuration(4, 4, 4)},
				{0, 4, 4, 4, 4.0, nil, duration.MakeDuration(5, 5, 5)},
				{0, 5, 5, 5, 5.0, 5.0, nil},
			},
			expected: tuples{
				{3.0, 3.0, 3.0, 3.0, 3.0, duration.MakeDuration(3, 3, 3)},
			},
			typs:    []*types.T{types.Int, types.Int2, types.Int4, types.Int, types.Decimal, types.Float, types.Interval},
			aggFns:  []execinfrapb.AggregatorSpec_Func{avgFn, avgFn, avgFn, avgFn, avgFn, avgFn},
			aggCols: [][]uint32{{1}, {2}, {3}, {4}, {5}, {6}},
			name:    "AVG on all types",
		},
		{
			input: tuples{
				{1, "1"},
				{1, "2"},
				{1, "3"},
				{2, nil},
				{2, "1"},
				{2, "2"},
				{3, "1"},
				{3, nil},
				{3, "2"},
				{4, nil},
				{4, nil},
			},
			expected: tuples{
				{"123"},
				{"12"},
				{"12"},
				{nil},
			},
			typs:      []*types.T{types.Int, types.Bytes},
			aggFns:    []execinfrapb.AggregatorSpec_Func{execinfrapb.AggregatorSpec_CONCAT_AGG},
			groupCols: []uint32{0},
			aggCols:   [][]uint32{{1}},
		},
	}

	evalCtx := tree.MakeTestingEvalContext(cluster.MakeTestingClusterSettings())
	defer evalCtx.Stop(context.Background())
	for _, agg := range aggTypes {
		for _, tc := range testCases {
			log.Infof(context.Background(), "%s/%s/Randomized", agg.name, tc.name)
			if err := tc.init(); err != nil {
				t.Fatal(err)
			}
			constructors, constArguments, outputTypes, err := ProcessAggregations(
				&evalCtx, nil /* semaCtx */, tc.spec.Aggregations, tc.typs,
			)
			require.NoError(t, err)
			runTestsWithTyps(t, []tuples{tc.input}, [][]*types.T{tc.typs}, tc.expected, unorderedVerifier,
				func(input []colexecbase.Operator) (colexecbase.Operator, error) {
					return agg.new(
						testAllocator, input[0], tc.typs, tc.spec, &evalCtx,
						constructors, constArguments, outputTypes, false, /* isScalar */
					)
				})
		}
	}
}

func TestAggregatorAllFunctions(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	testCases := []aggregatorTestCase{
		{
			aggFns: []execinfrapb.AggregatorSpec_Func{
				execinfrapb.AggregatorSpec_ANY_NOT_NULL,
				execinfrapb.AggregatorSpec_ANY_NOT_NULL,
				execinfrapb.AggregatorSpec_AVG,
				execinfrapb.AggregatorSpec_COUNT_ROWS,
				execinfrapb.AggregatorSpec_COUNT,
				execinfrapb.AggregatorSpec_SUM,
				execinfrapb.AggregatorSpec_SUM_INT,
				execinfrapb.AggregatorSpec_MIN,
				execinfrapb.AggregatorSpec_MAX,
				execinfrapb.AggregatorSpec_BOOL_AND,
				execinfrapb.AggregatorSpec_BOOL_OR,
				execinfrapb.AggregatorSpec_CONCAT_AGG,
			},
			aggCols: [][]uint32{{0}, {4}, {1}, {}, {1}, {1}, {2}, {2}, {2}, {3}, {3}, {4}},
			typs:    []*types.T{types.Int, types.Decimal, types.Int, types.Bool, types.Bytes},
			input: tuples{
				{0, 3.1, 2, true, "zero"},
				{0, 1.1, 3, false, "zero"},
				{1, 1.1, 1, false, "one"},
				{1, 4.1, 0, false, "one"},
				{2, 1.1, 1, true, "two"},
				{3, 4.1, 0, false, "three"},
				{3, 5.1, 0, true, "three"},
			},
			expected: tuples{
				{0, "zero", 2.1, 2, 2, 4.2, 5, 2, 3, false, true, "zerozero"},
				{1, "one", 2.6, 2, 2, 5.2, 1, 0, 1, false, false, "oneone"},
				{2, "two", 1.1, 1, 1, 1.1, 1, 1, 1, true, true, "two"},
				{3, "three", 4.6, 2, 2, 9.2, 0, 0, 0, false, true, "threethree"},
			},
			convToDecimal: true,
		},

		// Test case for null handling.
		{
			aggFns: []execinfrapb.AggregatorSpec_Func{
				execinfrapb.AggregatorSpec_ANY_NOT_NULL,
				execinfrapb.AggregatorSpec_ANY_NOT_NULL,
				execinfrapb.AggregatorSpec_COUNT_ROWS,
				execinfrapb.AggregatorSpec_COUNT,
				execinfrapb.AggregatorSpec_SUM,
				execinfrapb.AggregatorSpec_SUM_INT,
				execinfrapb.AggregatorSpec_MIN,
				execinfrapb.AggregatorSpec_MAX,
				execinfrapb.AggregatorSpec_AVG,
				execinfrapb.AggregatorSpec_BOOL_AND,
				execinfrapb.AggregatorSpec_BOOL_OR,
				execinfrapb.AggregatorSpec_CONCAT_AGG,
			},
			aggCols: [][]uint32{{0}, {1}, {}, {1}, {1}, {2}, {2}, {2}, {1}, {3}, {3}, {4}},
			typs:    []*types.T{types.Int, types.Decimal, types.Int, types.Bool, types.Bytes},
			input: tuples{
				{nil, 1.1, 4, true, "a"},
				{0, nil, nil, nil, nil},
				{0, 3.1, 5, nil, "b"},
				{1, nil, nil, nil, nil},
				{1, nil, nil, false, nil},
			},
			expected: tuples{
				{nil, 1.1, 1, 1, 1.1, 4, 4, 4, 1.1, true, true, "a"},
				{0, 3.1, 2, 1, 3.1, 5, 5, 5, 3.1, nil, nil, "b"},
				{1, nil, 2, 0, nil, nil, nil, nil, nil, false, false, nil},
			},
			convToDecimal: true,
		},
	}

	evalCtx := tree.MakeTestingEvalContext(cluster.MakeTestingClusterSettings())
	defer evalCtx.Stop(context.Background())
	for _, agg := range aggTypes {
		for i, tc := range testCases {
			log.Infof(context.Background(), "%s/%d", agg.name, i)
			if err := tc.init(); err != nil {
				t.Fatal(err)
			}
			constructors, constArguments, outputTypes, err := ProcessAggregations(
				&evalCtx, nil /* semaCtx */, tc.spec.Aggregations, tc.typs,
			)
			require.NoError(t, err)
			verifier := orderedVerifier
			if strings.Contains(agg.name, "hash") {
				verifier = unorderedVerifier
			}
			runTestsWithTyps(
				t,
				[]tuples{tc.input},
				[][]*types.T{tc.typs},
				tc.expected,
				verifier,
				func(input []colexecbase.Operator) (colexecbase.Operator, error) {
					return agg.new(
						testAllocator, input[0], tc.typs, tc.spec, &evalCtx,
						constructors, constArguments, outputTypes, false, /* isScalar */
					)
				})
		}
	}
}

func TestAggregatorRandom(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)

	evalCtx := tree.MakeTestingEvalContext(cluster.MakeTestingClusterSettings())
	defer evalCtx.Stop(context.Background())
	// This test aggregates random inputs, keeping track of the expected results
	// to make sure the aggregations are correct.
	rng, _ := randutil.NewPseudoRand()
	for _, groupSize := range []int{1, 2, coldata.BatchSize() / 4, coldata.BatchSize() / 2} {
		if groupSize == 0 {
			// We might be varying coldata.BatchSize() so that when it is divided by
			// 4, groupSize is 0. We want to skip such configuration.
			continue
		}
		for _, numInputBatches := range []int{1, 2, 64} {
			for _, hasNulls := range []bool{true, false} {
				for _, agg := range aggTypes {
					log.Infof(context.Background(), "%s/groupSize=%d/numInputBatches=%d/hasNulls=%t", agg.name, groupSize, numInputBatches, hasNulls)
					nTuples := coldata.BatchSize() * numInputBatches
					typs := []*types.T{types.Int, types.Float}
					cols := []coldata.Vec{
						testAllocator.NewMemColumn(typs[0], nTuples),
						testAllocator.NewMemColumn(typs[1], nTuples),
					}
					groups, aggCol, aggColNulls := cols[0].Int64(), cols[1].Float64(), cols[1].Nulls()
					expectedTuples := tuples{}

					var expRowCounts, expCounts []int64
					var expSums, expMins, expMaxs []float64
					// SUM, MIN, MAX, and AVG aggregators can output null.
					var expNulls []bool
					curGroup := -1
					for i := range groups {
						if i%groupSize == 0 {
							if curGroup != -1 {
								if expNulls[curGroup] {
									expectedTuples = append(expectedTuples, tuple{
										expRowCounts[curGroup], expCounts[curGroup], nil, nil, nil, nil,
									})
								} else {
									expectedTuples = append(expectedTuples, tuple{
										expRowCounts[curGroup], expCounts[curGroup], expSums[curGroup], expMins[curGroup], expMaxs[curGroup], expSums[curGroup] / float64(expCounts[curGroup]),
									})
								}
							}
							expRowCounts = append(expRowCounts, 0)
							expCounts = append(expCounts, 0)
							expSums = append(expSums, 0)
							expMins = append(expMins, 2048)
							expMaxs = append(expMaxs, -2048)
							expNulls = append(expNulls, true)
							curGroup++
						}
						// Keep the inputs small so they are a realistic size. Using a
						// large range is not realistic and makes decimal operations
						// slower.
						aggCol[i] = 2048 * (rng.Float64() - 0.5)

						// NULL values contribute to the row count, so we're updating
						// the row counts outside of the if block.
						expRowCounts[curGroup]++
						if hasNulls && rng.Float64() < nullProbability {
							aggColNulls.SetNull(i)
						} else {
							expNulls[curGroup] = false
							expCounts[curGroup]++
							expSums[curGroup] += aggCol[i]
							expMins[curGroup] = min64(aggCol[i], expMins[curGroup])
							expMaxs[curGroup] = max64(aggCol[i], expMaxs[curGroup])
						}
						groups[i] = int64(curGroup)
					}
					// Add result for last group.
					if expNulls[curGroup] {
						expectedTuples = append(expectedTuples, tuple{
							expRowCounts[curGroup], expCounts[curGroup], nil, nil, nil, nil,
						})
					} else {
						expectedTuples = append(expectedTuples, tuple{
							expRowCounts[curGroup], expCounts[curGroup], expSums[curGroup], expMins[curGroup], expMaxs[curGroup], expSums[curGroup] / float64(expCounts[curGroup]),
						})
					}

					source := newChunkingBatchSource(typs, cols, nTuples)
					tc := aggregatorTestCase{
						typs: typs,
						aggFns: []execinfrapb.AggregatorSpec_Func{
							execinfrapb.AggregatorSpec_COUNT_ROWS,
							execinfrapb.AggregatorSpec_COUNT,
							execinfrapb.AggregatorSpec_SUM,
							execinfrapb.AggregatorSpec_MIN,
							execinfrapb.AggregatorSpec_MAX,
							execinfrapb.AggregatorSpec_AVG,
						},
						groupCols: []uint32{0},
						aggCols:   [][]uint32{{}, {1}, {1}, {1}, {1}, {1}},
					}
					require.NoError(t, tc.init())
					constructors, constArguments, outputTypes, err := ProcessAggregations(
						&evalCtx, nil /* semaCtx */, tc.spec.Aggregations, tc.typs,
					)
					require.NoError(t, err)
					a, err := agg.new(
						testAllocator, source, tc.typs, tc.spec, &evalCtx,
						constructors, constArguments, outputTypes, false, /* isScalar */
					)
					if err != nil {
						t.Fatal(err)
					}
					a.Init()

					testOutput := newOpTestOutput(a, expectedTuples)
					if strings.Contains(agg.name, "hash") {
						err = testOutput.VerifyAnyOrder()
					} else {
						err = testOutput.Verify()
					}

					if err != nil {
						t.Fatal(err)
					}
				}
			}
		}
	}
}

func benchmarkAggregateFunction(
	b *testing.B,
	agg aggType,
	aggFn execinfrapb.AggregatorSpec_Func,
	aggInputTypes []*types.T,
	groupSize int,
	nullProb float64,
) {
	rng, _ := randutil.NewPseudoRand()
	ctx := context.Background()
	evalCtx := tree.MakeTestingEvalContext(cluster.MakeTestingClusterSettings())
	defer evalCtx.Stop(ctx)
	aggMemAcc := evalCtx.Mon.MakeBoundAccount()
	defer aggMemAcc.Close(ctx)
	evalCtx.SingleDatumAggMemAccount = &aggMemAcc
	const numInputBatches = 64
	const bytesFixedLength = 8
	typs := append([]*types.T{types.Int}, aggInputTypes...)
	nTuples := numInputBatches * coldata.BatchSize()
	cols := make([]coldata.Vec, len(typs))
	for i := range typs {
		cols[i] = testAllocator.NewMemColumn(typs[i], nTuples)
	}
	groups := cols[0].Int64()
	curGroup := -1
	for i := 0; i < nTuples; i++ {
		if groupSize == 1 || i%groupSize == 0 {
			curGroup++
		}
		groups[i] = int64(curGroup)
	}
	for _, col := range cols[1:] {
		coldatatestutils.RandomVec(coldatatestutils.RandomVecArgs{
			Rand:             rng,
			Vec:              col,
			N:                nTuples,
			NullProbability:  nullProb,
			BytesFixedLength: bytesFixedLength,
		})
	}
	if aggFn == execinfrapb.AggregatorSpec_SUM_INT {
		// Integer summation of random Int64 values can lead
		// to overflow, and we will panic. To go around it, we
		// restrict the range of values.
		vals := cols[1].Int64()
		for i := range vals {
			vals[i] = vals[i] % 1024
		}
	}
	source := newChunkingBatchSource(typs, cols, nTuples)

	aggCols := make([]uint32, len(aggInputTypes))
	for i := range aggCols {
		aggCols[i] = uint32(i + 1)
	}
	tc := aggregatorTestCase{
		typs:      typs,
		aggFns:    []execinfrapb.AggregatorSpec_Func{aggFn},
		groupCols: []uint32{0},
		aggCols:   [][]uint32{aggCols},
	}
	require.NoError(b, tc.init())
	constructors, constArguments, outputTypes, err := ProcessAggregations(
		&evalCtx, nil /* semaCtx */, tc.spec.Aggregations, tc.typs,
	)
	require.NoError(b, err)
	a, err := agg.new(
		testAllocator, source, typs, tc.spec, &evalCtx,
		constructors, constArguments, outputTypes, false, /* isScalar */
	)
	if err != nil {
		skip.IgnoreLint(b)
	}
	a.Init()

	b.ResetTimer()

	fName := execinfrapb.AggregatorSpec_Func_name[int32(aggFn)]
	// Only count the aggregation columns.
	var argumentsSize int
	if len(aggInputTypes) > 0 {
		for _, typ := range aggInputTypes {
			if typ.Identical(types.Bool) {
				argumentsSize++
			} else {
				argumentsSize += 8
			}
		}
	} else {
		// For COUNT_ROWS we'll just use 8 bytes.
		argumentsSize = 8
	}
	b.Run(fmt.Sprintf(
		"%s/%s/%s/groupSize=%d/hasNulls=%t/numInputBatches=%d",
		fName, agg.name, aggInputTypes, groupSize, nullProb > 0, numInputBatches),
		func(b *testing.B) {
			b.SetBytes(int64(argumentsSize * len(aggInputTypes) * nTuples))
			for i := 0; i < b.N; i++ {
				a.(resetter).reset(ctx)
				// Exhaust aggregator until all batches have been read.
				for b := a.Next(ctx); b.Length() != 0; b = a.Next(ctx) {
				}
			}
			require.NoError(b, a.(Closer).Close(ctx))
		},
	)
}

// BenchmarkAggregator runs the benchmark both aggregators with diverse data
// source parameters but using a single aggregate function. The goal of this
// benchmark is measuring the performance of the aggregators themselves
// depending on the parameters of the input.
func BenchmarkAggregator(b *testing.B) {
	aggFn := execinfrapb.AggregatorSpec_MIN
	for _, agg := range aggTypes {
		for _, groupSize := range []int{1, 2, 32, 128, coldata.BatchSize() / 2, coldata.BatchSize()} {
			for _, nullProb := range []float64{0.0, nullProbability} {
				benchmarkAggregateFunction(b, agg, aggFn, []*types.T{types.Int}, groupSize, nullProb)
			}
		}
	}
}

// BenchmarkAllOptimizedAggregateFunctions runs the benchmark of all optimized
// aggregate functions in 4 configurations (hash vs ordered, and small groups
// vs big groups). Such configurations were chosen since they provide good
// enough signal on the speeds of aggregate functions. For more diverse
// configurations look at BenchmarkAggregator.
func BenchmarkAllOptimizedAggregateFunctions(b *testing.B) {
	for aggFnNumber := 0; aggFnNumber < len(execinfrapb.AggregatorSpec_Func_name); aggFnNumber++ {
		aggFn := execinfrapb.AggregatorSpec_Func(aggFnNumber)
		if !isAggOptimized(aggFn) {
			continue
		}
		for _, agg := range aggTypes {
			var aggInputTypes []*types.T
			switch aggFn {
			case execinfrapb.AggregatorSpec_BOOL_AND, execinfrapb.AggregatorSpec_BOOL_OR:
				aggInputTypes = []*types.T{types.Bool}
			case execinfrapb.AggregatorSpec_CONCAT_AGG:
				aggInputTypes = []*types.T{types.Bytes}
			case execinfrapb.AggregatorSpec_COUNT_ROWS:
			default:
				aggInputTypes = []*types.T{types.Int}
			}
			for _, groupSize := range []int{1, coldata.BatchSize()} {
				benchmarkAggregateFunction(b, agg, aggFn, aggInputTypes, groupSize, nullProbability)
			}
		}
	}
}

func TestHashAggregator(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	tcs := []aggregatorTestCase{
		{
			// Test carry between output batches.
			input: tuples{
				{0, 1},
				{1, 5},
				{0, 4},
				{0, 2},
				{2, 6},
				{0, 3},
				{0, 7},
			},
			typs:      []*types.T{types.Int, types.Int},
			groupCols: []uint32{0},
			aggCols:   [][]uint32{{1}},

			expected: tuples{
				{5},
				{6},
				{17},
			},

			name: "carryBetweenBatches",
		},
		{
			// Test a single row input source.
			input: tuples{
				{5},
			},
			typs:      []*types.T{types.Int},
			groupCols: []uint32{0},
			aggCols:   [][]uint32{{0}},

			expected: tuples{
				{5},
			},

			name: "singleRowInput",
		},
		{
			// Test bucket collisions.
			input: tuples{
				{0, 3},
				{0, 4},
				{HashTableNumBuckets, 6},
				{0, 5},
				{HashTableNumBuckets, 7},
			},
			typs:      []*types.T{types.Int, types.Int},
			groupCols: []uint32{0},
			aggCols:   [][]uint32{{1}},

			expected: tuples{
				{12},
				{13},
			},

			name: "bucketCollision",
		},
		{
			input: tuples{
				{0, 1, 1.3},
				{0, 1, 1.6},
				{0, 1, 0.5},
				{1, 1, 1.2},
			},
			typs:          []*types.T{types.Int, types.Int, types.Decimal},
			convToDecimal: true,

			aggFns:    []execinfrapb.AggregatorSpec_Func{execinfrapb.AggregatorSpec_SUM, execinfrapb.AggregatorSpec_SUM_INT},
			groupCols: []uint32{0, 1},
			aggCols: [][]uint32{
				{2}, {1},
			},

			expected: tuples{
				{3.4, 3},
				{1.2, 1},
			},

			name: "decimalSums",
		},
		{
			// Test unused input columns.
			input: tuples{
				{0, 1, 2, 3},
				{0, 1, 4, 5},
				{1, 1, 3, 7},
				{1, 2, 4, 9},
				{0, 1, 6, 11},
				{1, 2, 6, 13},
			},
			typs:      []*types.T{types.Int, types.Int, types.Int, types.Int},
			groupCols: []uint32{0, 1},
			aggCols:   [][]uint32{{3}},

			expected: tuples{
				{7},
				{19},
				{22},
			},

			name: "unusedInputCol",
		},
	}

	evalCtx := tree.MakeTestingEvalContext(cluster.MakeTestingClusterSettings())
	defer evalCtx.Stop(context.Background())
	for _, numOfHashBuckets := range []int{0 /* no limit */, 1, coldata.BatchSize()} {
		for _, tc := range tcs {
			if err := tc.init(); err != nil {
				t.Fatal(err)
			}
			constructors, constArguments, outputTypes, err := ProcessAggregations(
				&evalCtx, nil /* semaCtx */, tc.spec.Aggregations, tc.typs,
			)
			require.NoError(t, err)
			log.Infof(context.Background(), "numOfHashBuckets=%d", numOfHashBuckets)
			runTests(t, []tuples{tc.input}, tc.expected, unorderedVerifier, func(sources []colexecbase.Operator) (colexecbase.Operator, error) {
				a, err := NewHashAggregator(
					testAllocator, sources[0], tc.typs, tc.spec, &evalCtx,
					constructors, constArguments, outputTypes,
				)
				a.(*hashAggregator).testingKnobs.numOfHashBuckets = uint64(numOfHashBuckets)
				return a, err
			})
		}
	}
}

func min64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

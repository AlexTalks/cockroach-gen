// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"unsafe"

	"github.com/cockroachdb/apd"
	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/col/typeconv"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase/colexecerror"
	"github.com/cockroachdb/cockroach/pkg/sql/colmem"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/errors"
)

func newAvgAgg(allocator *colmem.Allocator, t *types.T) (aggregateFunc, error) {
	switch typeconv.TypeFamilyToCanonicalTypeFamily(t.Family()) {
	case types.DecimalFamily:
		switch t.Width() {
		case -1:
		default:
			allocator.AdjustMemoryUsage(int64(sizeOfAvgDecimalAgg))
			return &avgDecimalAgg{}, nil
		}
	case types.FloatFamily:
		switch t.Width() {
		case -1:
		default:
			allocator.AdjustMemoryUsage(int64(sizeOfAvgFloat64Agg))
			return &avgFloat64Agg{}, nil
		}
	}
	return nil, errors.Errorf("unsupported avg agg type %s", t.Name())
}

type avgDecimalAgg struct {
	groups  []bool
	scratch struct {
		curIdx int
		// curSum keeps track of the sum of elements belonging to the current group,
		// so we can index into the slice once per group, instead of on each
		// iteration.
		curSum apd.Decimal
		// curCount keeps track of the number of elements that we've seen
		// belonging to the current group.
		curCount int64
		// vec points to the output vector.
		vec []apd.Decimal
		// nulls points to the output null vector that we are updating.
		nulls *coldata.Nulls
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
}

var _ aggregateFunc = &avgDecimalAgg{}

const sizeOfAvgDecimalAgg = unsafe.Sizeof(&avgDecimalAgg{})

func (a *avgDecimalAgg) Init(groups []bool, v coldata.Vec) {
	a.groups = groups
	a.scratch.vec = v.Decimal()
	a.scratch.nulls = v.Nulls()
	a.Reset()
}

func (a *avgDecimalAgg) Reset() {
	a.scratch.curIdx = -1
	a.scratch.curSum = zeroDecimalValue
	a.scratch.curCount = 0
	a.scratch.foundNonNullForCurrentGroup = false
	a.scratch.nulls.UnsetNulls()
}

func (a *avgDecimalAgg) CurrentOutputIndex() int {
	return a.scratch.curIdx
}

func (a *avgDecimalAgg) SetOutputIndex(idx int) {
	if a.scratch.curIdx != -1 {
		a.scratch.curIdx = idx
		a.scratch.nulls.UnsetNullsAfter(idx + 1)
	}
}

func (a *avgDecimalAgg) Compute(b coldata.Batch, inputIdxs []uint32) {
	inputLen := b.Length()
	vec, sel := b.ColVec(int(inputIdxs[0])), b.Selection()
	col, nulls := vec.Decimal(), vec.Nulls()
	if nulls.MaybeHasNulls() {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {

							a.scratch.vec[a.scratch.curIdx].SetInt64(a.scratch.curCount)
							if _, err := tree.DecimalCtx.Quo(&a.scratch.vec[a.scratch.curIdx], &a.scratch.curSum, &a.scratch.vec[a.scratch.curIdx]); err != nil {
								colexecerror.InternalError(err)
							}
						}
					}
					a.scratch.curIdx++
					a.scratch.curSum = zeroDecimalValue
					a.scratch.curCount = 0

					a.scratch.foundNonNullForCurrentGroup = false
				}
				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {
					if _, err := tree.ExactCtx.Add(&a.scratch.curSum, &a.scratch.curSum, &col[i]); err != nil {
						colexecerror.ExpectedError(err)
					}
					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {

							a.scratch.vec[a.scratch.curIdx].SetInt64(a.scratch.curCount)
							if _, err := tree.DecimalCtx.Quo(&a.scratch.vec[a.scratch.curIdx], &a.scratch.curSum, &a.scratch.vec[a.scratch.curIdx]); err != nil {
								colexecerror.InternalError(err)
							}
						}
					}
					a.scratch.curIdx++
					a.scratch.curSum = zeroDecimalValue
					a.scratch.curCount = 0

					a.scratch.foundNonNullForCurrentGroup = false
				}
				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {
					if _, err := tree.ExactCtx.Add(&a.scratch.curSum, &a.scratch.curSum, &col[i]); err != nil {
						colexecerror.ExpectedError(err)
					}
					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	} else {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {

							a.scratch.vec[a.scratch.curIdx].SetInt64(a.scratch.curCount)
							if _, err := tree.DecimalCtx.Quo(&a.scratch.vec[a.scratch.curIdx], &a.scratch.curSum, &a.scratch.vec[a.scratch.curIdx]); err != nil {
								colexecerror.InternalError(err)
							}
						}
					}
					a.scratch.curIdx++
					a.scratch.curSum = zeroDecimalValue
					a.scratch.curCount = 0

				}
				var isNull bool
				isNull = false
				if !isNull {
					if _, err := tree.ExactCtx.Add(&a.scratch.curSum, &a.scratch.curSum, &col[i]); err != nil {
						colexecerror.ExpectedError(err)
					}
					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {

							a.scratch.vec[a.scratch.curIdx].SetInt64(a.scratch.curCount)
							if _, err := tree.DecimalCtx.Quo(&a.scratch.vec[a.scratch.curIdx], &a.scratch.curSum, &a.scratch.vec[a.scratch.curIdx]); err != nil {
								colexecerror.InternalError(err)
							}
						}
					}
					a.scratch.curIdx++
					a.scratch.curSum = zeroDecimalValue
					a.scratch.curCount = 0

				}
				var isNull bool
				isNull = false
				if !isNull {
					if _, err := tree.ExactCtx.Add(&a.scratch.curSum, &a.scratch.curSum, &col[i]); err != nil {
						colexecerror.ExpectedError(err)
					}
					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *avgDecimalAgg) Flush() {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// NULL.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.scratch.nulls.SetNull(a.scratch.curIdx)
	} else {

		a.scratch.vec[a.scratch.curIdx].SetInt64(a.scratch.curCount)
		if _, err := tree.DecimalCtx.Quo(&a.scratch.vec[a.scratch.curIdx], &a.scratch.curSum, &a.scratch.vec[a.scratch.curIdx]); err != nil {
			colexecerror.InternalError(err)
		}
	}
	a.scratch.curIdx++
}

func (a *avgDecimalAgg) HandleEmptyInputScalar() {
	a.scratch.nulls.SetNull(0)
}

type avgFloat64Agg struct {
	groups  []bool
	scratch struct {
		curIdx int
		// curSum keeps track of the sum of elements belonging to the current group,
		// so we can index into the slice once per group, instead of on each
		// iteration.
		curSum float64
		// curCount keeps track of the number of elements that we've seen
		// belonging to the current group.
		curCount int64
		// vec points to the output vector.
		vec []float64
		// nulls points to the output null vector that we are updating.
		nulls *coldata.Nulls
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
}

var _ aggregateFunc = &avgFloat64Agg{}

const sizeOfAvgFloat64Agg = unsafe.Sizeof(&avgFloat64Agg{})

func (a *avgFloat64Agg) Init(groups []bool, v coldata.Vec) {
	a.groups = groups
	a.scratch.vec = v.Float64()
	a.scratch.nulls = v.Nulls()
	a.Reset()
}

func (a *avgFloat64Agg) Reset() {
	a.scratch.curIdx = -1
	a.scratch.curSum = zeroFloat64Value
	a.scratch.curCount = 0
	a.scratch.foundNonNullForCurrentGroup = false
	a.scratch.nulls.UnsetNulls()
}

func (a *avgFloat64Agg) CurrentOutputIndex() int {
	return a.scratch.curIdx
}

func (a *avgFloat64Agg) SetOutputIndex(idx int) {
	if a.scratch.curIdx != -1 {
		a.scratch.curIdx = idx
		a.scratch.nulls.UnsetNullsAfter(idx + 1)
	}
}

func (a *avgFloat64Agg) Compute(b coldata.Batch, inputIdxs []uint32) {
	inputLen := b.Length()
	vec, sel := b.ColVec(int(inputIdxs[0])), b.Selection()
	col, nulls := vec.Float64(), vec.Nulls()
	if nulls.MaybeHasNulls() {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curSum / float64(a.scratch.curCount)
						}
					}
					a.scratch.curIdx++
					a.scratch.curSum = zeroFloat64Value
					a.scratch.curCount = 0

					a.scratch.foundNonNullForCurrentGroup = false
				}
				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {
					a.scratch.curSum = float64(a.scratch.curSum) + float64(col[i])
					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curSum / float64(a.scratch.curCount)
						}
					}
					a.scratch.curIdx++
					a.scratch.curSum = zeroFloat64Value
					a.scratch.curCount = 0

					a.scratch.foundNonNullForCurrentGroup = false
				}
				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {
					a.scratch.curSum = float64(a.scratch.curSum) + float64(col[i])
					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	} else {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curSum / float64(a.scratch.curCount)
						}
					}
					a.scratch.curIdx++
					a.scratch.curSum = zeroFloat64Value
					a.scratch.curCount = 0

				}
				var isNull bool
				isNull = false
				if !isNull {
					a.scratch.curSum = float64(a.scratch.curSum) + float64(col[i])
					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curSum / float64(a.scratch.curCount)
						}
					}
					a.scratch.curIdx++
					a.scratch.curSum = zeroFloat64Value
					a.scratch.curCount = 0

				}
				var isNull bool
				isNull = false
				if !isNull {
					a.scratch.curSum = float64(a.scratch.curSum) + float64(col[i])
					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *avgFloat64Agg) Flush() {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// NULL.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.scratch.nulls.SetNull(a.scratch.curIdx)
	} else {
		a.scratch.vec[a.scratch.curIdx] = a.scratch.curSum / float64(a.scratch.curCount)
	}
	a.scratch.curIdx++
}

func (a *avgFloat64Agg) HandleEmptyInputScalar() {
	a.scratch.nulls.SetNull(0)
}

// Code generated by execgen; DO NOT EDIT.
// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

//

package colexec

import (
	"context"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase/colexecerror"
	"github.com/cockroachdb/cockroach/pkg/sql/colmem"
	"github.com/cockroachdb/cockroach/pkg/sql/execinfrapb"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/errors"
)

// Remove unused warning.
var _ = colexecerror.InternalError

// TODO(yuzefovich): add benchmarks.

// NewRankOperator creates a new Operator that computes window functions RANK
// or DENSE_RANK (depending on the passed in windowFn).
// outputColIdx specifies in which coldata.Vec the operator should put its
// output (if there is no such column, a new column is appended).
func NewRankOperator(
	allocator *colmem.Allocator,
	input colexecbase.Operator,
	windowFn execinfrapb.WindowerSpec_WindowFunc,
	orderingCols []execinfrapb.Ordering_Column,
	outputColIdx int,
	partitionColIdx int,
	peersColIdx int,
) (colexecbase.Operator, error) {
	if len(orderingCols) == 0 {
		return NewConstOp(allocator, input, types.Int, int64(1), outputColIdx)
	}
	input = newVectorTypeEnforcer(allocator, input, types.Int, outputColIdx)
	initFields := rankInitFields{
		OneInputNode:    NewOneInputNode(input),
		allocator:       allocator,
		outputColIdx:    outputColIdx,
		partitionColIdx: partitionColIdx,
		peersColIdx:     peersColIdx,
	}
	switch windowFn {
	case execinfrapb.WindowerSpec_RANK:
		if partitionColIdx != columnOmitted {
			return &rankWithPartitionOp{rankInitFields: initFields}, nil
		}
		return &rankNoPartitionOp{rankInitFields: initFields}, nil
	case execinfrapb.WindowerSpec_DENSE_RANK:
		if partitionColIdx != columnOmitted {
			return &denseRankWithPartitionOp{rankInitFields: initFields}, nil
		}
		return &denseRankNoPartitionOp{rankInitFields: initFields}, nil
	default:
		return nil, errors.Errorf("unsupported rank type %s", windowFn)
	}
}

//

type rankInitFields struct {
	OneInputNode

	allocator       *colmem.Allocator
	outputColIdx    int
	partitionColIdx int
	peersColIdx     int
}

//
//
//

//

type rankNoPartitionOp struct {
	rankInitFields

	// rank indicates which rank should be assigned to the next tuple.
	rank int64
	// rankIncrement indicates by how much rank should be incremented when a
	// tuple distinct from the previous one on the ordering columns is seen.
	rankIncrement int64
}

var _ colexecbase.Operator = &rankNoPartitionOp{}

func (r *rankNoPartitionOp) Init() {
	r.Input().Init()
	// All rank functions start counting from 1. Before we assign the rank to a
	// tuple in the batch, we first increment r.rank, so setting this
	// rankIncrement to 1 will update r.rank to 1 on the very first tuple (as
	// desired).
	r.rankIncrement = 1
}

func (r *rankNoPartitionOp) Next(ctx context.Context) coldata.Batch {
	batch := r.Input().Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	//
	peersCol := batch.ColVec(r.peersColIdx).Bool()
	rankVec := batch.ColVec(r.outputColIdx)
	if rankVec.MaybeHasNulls() {
		// We need to make sure that there are no left over null values in the
		// output vector.
		rankVec.Nulls().UnsetNulls()
	}
	rankCol := rankVec.Int64()
	sel := batch.Selection()
	if sel != nil {
		for _, i := range sel[:n] {
			//
			if peersCol[i] {
				r.rank += r.rankIncrement
				r.rankIncrement = 1
				rankCol[i] = r.rank
			} else {
				rankCol[i] = r.rank
				r.rankIncrement++
			}
			//
		}
	} else {
		for i := range rankCol[:n] {
			//
			if peersCol[i] {
				r.rank += r.rankIncrement
				r.rankIncrement = 1
				rankCol[i] = r.rank
			} else {
				rankCol[i] = r.rank
				r.rankIncrement++
			}
			//
		}
	}
	return batch
}

//

type rankWithPartitionOp struct {
	rankInitFields

	// rank indicates which rank should be assigned to the next tuple.
	rank int64
	// rankIncrement indicates by how much rank should be incremented when a
	// tuple distinct from the previous one on the ordering columns is seen.
	rankIncrement int64
}

var _ colexecbase.Operator = &rankWithPartitionOp{}

func (r *rankWithPartitionOp) Init() {
	r.Input().Init()
	// All rank functions start counting from 1. Before we assign the rank to a
	// tuple in the batch, we first increment r.rank, so setting this
	// rankIncrement to 1 will update r.rank to 1 on the very first tuple (as
	// desired).
	r.rankIncrement = 1
}

func (r *rankWithPartitionOp) Next(ctx context.Context) coldata.Batch {
	batch := r.Input().Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	//
	partitionCol := batch.ColVec(r.partitionColIdx).Bool()
	//
	peersCol := batch.ColVec(r.peersColIdx).Bool()
	rankVec := batch.ColVec(r.outputColIdx)
	if rankVec.MaybeHasNulls() {
		// We need to make sure that there are no left over null values in the
		// output vector.
		rankVec.Nulls().UnsetNulls()
	}
	rankCol := rankVec.Int64()
	sel := batch.Selection()
	if sel != nil {
		for _, i := range sel[:n] {
			//
			if partitionCol[i] {
				// We need to reset the internal state because of the new partition.
				// Note that the beginning of new partition necessarily starts a new
				// peer group, so peersCol[i] *must* be true, and we will correctly
				// update the rank before setting it to rankCol.
				r.rank = 0
				r.rankIncrement = 1
			}
			//
			if peersCol[i] {
				r.rank += r.rankIncrement
				r.rankIncrement = 1
				rankCol[i] = r.rank
			} else {
				rankCol[i] = r.rank
				r.rankIncrement++
			}
			//
		}
	} else {
		for i := range rankCol[:n] {
			//
			if partitionCol[i] {
				// We need to reset the internal state because of the new partition.
				// Note that the beginning of new partition necessarily starts a new
				// peer group, so peersCol[i] *must* be true, and we will correctly
				// update the rank before setting it to rankCol.
				r.rank = 0
				r.rankIncrement = 1
			}
			//
			if peersCol[i] {
				r.rank += r.rankIncrement
				r.rankIncrement = 1
				rankCol[i] = r.rank
			} else {
				rankCol[i] = r.rank
				r.rankIncrement++
			}
			//
		}
	}
	return batch
}

//

type denseRankNoPartitionOp struct {
	rankInitFields

	// rank indicates which rank should be assigned to the next tuple.
	rank int64
	// rankIncrement indicates by how much rank should be incremented when a
	// tuple distinct from the previous one on the ordering columns is seen.
	rankIncrement int64
}

var _ colexecbase.Operator = &denseRankNoPartitionOp{}

func (r *denseRankNoPartitionOp) Init() {
	r.Input().Init()
	// All rank functions start counting from 1. Before we assign the rank to a
	// tuple in the batch, we first increment r.rank, so setting this
	// rankIncrement to 1 will update r.rank to 1 on the very first tuple (as
	// desired).
	r.rankIncrement = 1
}

func (r *denseRankNoPartitionOp) Next(ctx context.Context) coldata.Batch {
	batch := r.Input().Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	//
	peersCol := batch.ColVec(r.peersColIdx).Bool()
	rankVec := batch.ColVec(r.outputColIdx)
	if rankVec.MaybeHasNulls() {
		// We need to make sure that there are no left over null values in the
		// output vector.
		rankVec.Nulls().UnsetNulls()
	}
	rankCol := rankVec.Int64()
	sel := batch.Selection()
	if sel != nil {
		for _, i := range sel[:n] {
			//
			if peersCol[i] {
				r.rank++
				rankCol[i] = r.rank
			} else {
				rankCol[i] = r.rank

			}
			//
		}
	} else {
		for i := range rankCol[:n] {
			//
			if peersCol[i] {
				r.rank++
				rankCol[i] = r.rank
			} else {
				rankCol[i] = r.rank

			}
			//
		}
	}
	return batch
}

//

type denseRankWithPartitionOp struct {
	rankInitFields

	// rank indicates which rank should be assigned to the next tuple.
	rank int64
	// rankIncrement indicates by how much rank should be incremented when a
	// tuple distinct from the previous one on the ordering columns is seen.
	rankIncrement int64
}

var _ colexecbase.Operator = &denseRankWithPartitionOp{}

func (r *denseRankWithPartitionOp) Init() {
	r.Input().Init()
	// All rank functions start counting from 1. Before we assign the rank to a
	// tuple in the batch, we first increment r.rank, so setting this
	// rankIncrement to 1 will update r.rank to 1 on the very first tuple (as
	// desired).
	r.rankIncrement = 1
}

func (r *denseRankWithPartitionOp) Next(ctx context.Context) coldata.Batch {
	batch := r.Input().Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	//
	partitionCol := batch.ColVec(r.partitionColIdx).Bool()
	//
	peersCol := batch.ColVec(r.peersColIdx).Bool()
	rankVec := batch.ColVec(r.outputColIdx)
	if rankVec.MaybeHasNulls() {
		// We need to make sure that there are no left over null values in the
		// output vector.
		rankVec.Nulls().UnsetNulls()
	}
	rankCol := rankVec.Int64()
	sel := batch.Selection()
	if sel != nil {
		for _, i := range sel[:n] {
			//
			if partitionCol[i] {
				// We need to reset the internal state because of the new partition.
				// Note that the beginning of new partition necessarily starts a new
				// peer group, so peersCol[i] *must* be true, and we will correctly
				// update the rank before setting it to rankCol.
				r.rank = 0
				r.rankIncrement = 1
			}
			//
			if peersCol[i] {
				r.rank++
				rankCol[i] = r.rank
			} else {
				rankCol[i] = r.rank

			}
			//
		}
	} else {
		for i := range rankCol[:n] {
			//
			if partitionCol[i] {
				// We need to reset the internal state because of the new partition.
				// Note that the beginning of new partition necessarily starts a new
				// peer group, so peersCol[i] *must* be true, and we will correctly
				// update the rank before setting it to rankCol.
				r.rank = 0
				r.rankIncrement = 1
			}
			//
			if peersCol[i] {
				r.rank++
				rankCol[i] = r.rank
			} else {
				rankCol[i] = r.rank

			}
			//
		}
	}
	return batch
}

//

// Code generated by execgen; DO NOT EDIT.
// Copyright 2020 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase/colexecerror"
)

// Remove unused warning.
var _ = colexecerror.InternalError

func newBoolAndAgg() aggregateFunc {
	return &boolAndAgg{}
}

type boolAndAgg struct {
	done       bool
	sawNonNull bool

	groups []bool
	vec    []bool

	nulls  *coldata.Nulls
	curIdx int
	curAgg bool
}

func (b *boolAndAgg) Init(groups []bool, vec coldata.Vec) {
	b.groups = groups
	b.vec = vec.Bool()
	b.nulls = vec.Nulls()
	b.Reset()
}

func (b *boolAndAgg) Reset() {
	b.curIdx = -1
	b.nulls.UnsetNulls()
	b.done = false
	// true indicates whether we are doing an AND aggregate or OR aggregate.
	// For bool_and the true is true and for bool_or the true is false.
	b.curAgg = true
}

func (b *boolAndAgg) CurrentOutputIndex() int {
	return b.curIdx
}

func (b *boolAndAgg) SetOutputIndex(idx int) {
	if b.curIdx != -1 {
		b.curIdx = idx
		b.nulls.UnsetNullsAfter(idx)
	}
}

func (b *boolAndAgg) Compute(batch coldata.Batch, inputIdxs []uint32) {
	if b.done {
		return
	}
	inputLen := batch.Length()
	if inputLen == 0 {
		if !b.sawNonNull {
			b.nulls.SetNull(b.curIdx)
		} else {
			b.vec[b.curIdx] = b.curAgg
		}
		b.curIdx++
		b.done = true
		return
	}
	vec, sel := batch.ColVec(int(inputIdxs[0])), batch.Selection()
	col, nulls := vec.Bool(), vec.Nulls()
	if sel != nil {
		sel = sel[:inputLen]
		for _, i := range sel {
			if b.groups[i] {
				if b.curIdx >= 0 {
					if !b.sawNonNull {
						b.nulls.SetNull(b.curIdx)
					} else {
						b.vec[b.curIdx] = b.curAgg
					}
				}
				b.curIdx++
				b.curAgg = true
				b.sawNonNull = false
			}
			isNull := nulls.NullAt(i)
			if !isNull {
				b.curAgg = b.curAgg && col[i]
				b.sawNonNull = true
			}

		}
	} else {
		col = col[:inputLen]
		for i := range col {
			if b.groups[i] {
				if b.curIdx >= 0 {
					if !b.sawNonNull {
						b.nulls.SetNull(b.curIdx)
					} else {
						b.vec[b.curIdx] = b.curAgg
					}
				}
				b.curIdx++
				b.curAgg = true
				b.sawNonNull = false
			}
			isNull := nulls.NullAt(i)
			if !isNull {
				b.curAgg = b.curAgg && col[i]
				b.sawNonNull = true
			}

		}
	}
}

func (b *boolAndAgg) HandleEmptyInputScalar() {
	b.nulls.SetNull(0)
}

func newBoolOrAgg() aggregateFunc {
	return &boolOrAgg{}
}

type boolOrAgg struct {
	done       bool
	sawNonNull bool

	groups []bool
	vec    []bool

	nulls  *coldata.Nulls
	curIdx int
	curAgg bool
}

func (b *boolOrAgg) Init(groups []bool, vec coldata.Vec) {
	b.groups = groups
	b.vec = vec.Bool()
	b.nulls = vec.Nulls()
	b.Reset()
}

func (b *boolOrAgg) Reset() {
	b.curIdx = -1
	b.nulls.UnsetNulls()
	b.done = false
	// false indicates whether we are doing an AND aggregate or OR aggregate.
	// For bool_and the false is true and for bool_or the false is false.
	b.curAgg = false
}

func (b *boolOrAgg) CurrentOutputIndex() int {
	return b.curIdx
}

func (b *boolOrAgg) SetOutputIndex(idx int) {
	if b.curIdx != -1 {
		b.curIdx = idx
		b.nulls.UnsetNullsAfter(idx)
	}
}

func (b *boolOrAgg) Compute(batch coldata.Batch, inputIdxs []uint32) {
	if b.done {
		return
	}
	inputLen := batch.Length()
	if inputLen == 0 {
		if !b.sawNonNull {
			b.nulls.SetNull(b.curIdx)
		} else {
			b.vec[b.curIdx] = b.curAgg
		}
		b.curIdx++
		b.done = true
		return
	}
	vec, sel := batch.ColVec(int(inputIdxs[0])), batch.Selection()
	col, nulls := vec.Bool(), vec.Nulls()
	if sel != nil {
		sel = sel[:inputLen]
		for _, i := range sel {
			if b.groups[i] {
				if b.curIdx >= 0 {
					if !b.sawNonNull {
						b.nulls.SetNull(b.curIdx)
					} else {
						b.vec[b.curIdx] = b.curAgg
					}
				}
				b.curIdx++
				b.curAgg = false
				b.sawNonNull = false
			}
			isNull := nulls.NullAt(i)
			if !isNull {
				b.curAgg = b.curAgg || col[i]
				b.sawNonNull = true
			}

		}
	} else {
		col = col[:inputLen]
		for i := range col {
			if b.groups[i] {
				if b.curIdx >= 0 {
					if !b.sawNonNull {
						b.nulls.SetNull(b.curIdx)
					} else {
						b.vec[b.curIdx] = b.curAgg
					}
				}
				b.curIdx++
				b.curAgg = false
				b.sawNonNull = false
			}
			isNull := nulls.NullAt(i)
			if !isNull {
				b.curAgg = b.curAgg || col[i]
				b.sawNonNull = true
			}

		}
	}
}

func (b *boolOrAgg) HandleEmptyInputScalar() {
	b.nulls.SetNull(0)
}

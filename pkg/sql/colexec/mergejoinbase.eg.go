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
	"bytes"
	"fmt"
	"math"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/col/coldataext"
	"github.com/cockroachdb/cockroach/pkg/col/typeconv"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase/colexecerror"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
)

//

// isBufferedGroupFinished checks to see whether or not the buffered group
// corresponding to input continues in batch.
func (o *mergeJoinBase) isBufferedGroupFinished(
	input *mergeJoinInput, batch coldata.Batch, rowIdx int,
) bool {
	if batch.Length() == 0 {
		return true
	}
	bufferedGroup := o.proberState.lBufferedGroup
	if input == &o.right {
		bufferedGroup = o.proberState.rBufferedGroup
	}
	tupleToLookAtIdx := rowIdx
	sel := batch.Selection()
	if sel != nil {
		tupleToLookAtIdx = sel[rowIdx]
	}

	// Check all equality columns in the first row of batch to make sure we're in
	// the same group.
	for _, colIdx := range input.eqCols[:len(input.eqCols)] {
		switch input.canonicalTypeFamilies[colIdx] {
		//
		case types.BoolFamily:
			switch input.sourceTypes[colIdx].Width() {
			//
			case -1:
			default:
				// We perform this null check on every equality column of the first
				// buffered tuple regardless of the join type since it is done only once
				// per batch. In some cases (like INNER join, or LEFT OUTER join with the
				// right side being an input) this check will always return false since
				// nulls couldn't be buffered up though.
				// TODO(yuzefovich): consider templating this.
				bufferedNull := bufferedGroup.firstTuple[colIdx].MaybeHasNulls() && bufferedGroup.firstTuple[colIdx].Nulls().NullAt(0)
				incomingNull := batch.ColVec(int(colIdx)).MaybeHasNulls() && batch.ColVec(int(colIdx)).Nulls().NullAt(tupleToLookAtIdx)
				if o.joinType.IsSetOpJoin() {
					if bufferedNull && incomingNull {
						// We have a NULL match, so move onto the next column.
						continue
					}
				}
				if bufferedNull || incomingNull {
					return true
				}
				bufferedCol := bufferedGroup.firstTuple[colIdx].Bool()
				prevVal := bufferedCol[0]
				col := batch.ColVec(int(colIdx)).Bool()
				curVal := col[tupleToLookAtIdx]
				var match bool

				{
					var cmpResult int

					if !prevVal && curVal {
						cmpResult = -1
					} else if prevVal && !curVal {
						cmpResult = 1
					} else {
						cmpResult = 0
					}

					match = cmpResult == 0
				}

				if !match {
					return true
				}
				//
			}
		//
		case types.BytesFamily:
			switch input.sourceTypes[colIdx].Width() {
			//
			case -1:
			default:
				// We perform this null check on every equality column of the first
				// buffered tuple regardless of the join type since it is done only once
				// per batch. In some cases (like INNER join, or LEFT OUTER join with the
				// right side being an input) this check will always return false since
				// nulls couldn't be buffered up though.
				// TODO(yuzefovich): consider templating this.
				bufferedNull := bufferedGroup.firstTuple[colIdx].MaybeHasNulls() && bufferedGroup.firstTuple[colIdx].Nulls().NullAt(0)
				incomingNull := batch.ColVec(int(colIdx)).MaybeHasNulls() && batch.ColVec(int(colIdx)).Nulls().NullAt(tupleToLookAtIdx)
				if o.joinType.IsSetOpJoin() {
					if bufferedNull && incomingNull {
						// We have a NULL match, so move onto the next column.
						continue
					}
				}
				if bufferedNull || incomingNull {
					return true
				}
				bufferedCol := bufferedGroup.firstTuple[colIdx].Bytes()
				prevVal := bufferedCol.Get(0)
				col := batch.ColVec(int(colIdx)).Bytes()
				curVal := col.Get(tupleToLookAtIdx)
				var match bool

				{
					var cmpResult int
					cmpResult = bytes.Compare(prevVal, curVal)
					match = cmpResult == 0
				}

				if !match {
					return true
				}
				//
			}
		//
		case types.DecimalFamily:
			switch input.sourceTypes[colIdx].Width() {
			//
			case -1:
			default:
				// We perform this null check on every equality column of the first
				// buffered tuple regardless of the join type since it is done only once
				// per batch. In some cases (like INNER join, or LEFT OUTER join with the
				// right side being an input) this check will always return false since
				// nulls couldn't be buffered up though.
				// TODO(yuzefovich): consider templating this.
				bufferedNull := bufferedGroup.firstTuple[colIdx].MaybeHasNulls() && bufferedGroup.firstTuple[colIdx].Nulls().NullAt(0)
				incomingNull := batch.ColVec(int(colIdx)).MaybeHasNulls() && batch.ColVec(int(colIdx)).Nulls().NullAt(tupleToLookAtIdx)
				if o.joinType.IsSetOpJoin() {
					if bufferedNull && incomingNull {
						// We have a NULL match, so move onto the next column.
						continue
					}
				}
				if bufferedNull || incomingNull {
					return true
				}
				bufferedCol := bufferedGroup.firstTuple[colIdx].Decimal()
				prevVal := bufferedCol[0]
				col := batch.ColVec(int(colIdx)).Decimal()
				curVal := col[tupleToLookAtIdx]
				var match bool

				{
					var cmpResult int
					cmpResult = tree.CompareDecimals(&prevVal, &curVal)
					match = cmpResult == 0
				}

				if !match {
					return true
				}
				//
			}
		//
		case types.IntFamily:
			switch input.sourceTypes[colIdx].Width() {
			//
			case 16:
				// We perform this null check on every equality column of the first
				// buffered tuple regardless of the join type since it is done only once
				// per batch. In some cases (like INNER join, or LEFT OUTER join with the
				// right side being an input) this check will always return false since
				// nulls couldn't be buffered up though.
				// TODO(yuzefovich): consider templating this.
				bufferedNull := bufferedGroup.firstTuple[colIdx].MaybeHasNulls() && bufferedGroup.firstTuple[colIdx].Nulls().NullAt(0)
				incomingNull := batch.ColVec(int(colIdx)).MaybeHasNulls() && batch.ColVec(int(colIdx)).Nulls().NullAt(tupleToLookAtIdx)
				if o.joinType.IsSetOpJoin() {
					if bufferedNull && incomingNull {
						// We have a NULL match, so move onto the next column.
						continue
					}
				}
				if bufferedNull || incomingNull {
					return true
				}
				bufferedCol := bufferedGroup.firstTuple[colIdx].Int16()
				prevVal := bufferedCol[0]
				col := batch.ColVec(int(colIdx)).Int16()
				curVal := col[tupleToLookAtIdx]
				var match bool

				{
					var cmpResult int

					{
						a, b := int64(prevVal), int64(curVal)
						if a < b {
							cmpResult = -1
						} else if a > b {
							cmpResult = 1
						} else {
							cmpResult = 0
						}
					}

					match = cmpResult == 0
				}

				if !match {
					return true
				}
				//
			case 32:
				// We perform this null check on every equality column of the first
				// buffered tuple regardless of the join type since it is done only once
				// per batch. In some cases (like INNER join, or LEFT OUTER join with the
				// right side being an input) this check will always return false since
				// nulls couldn't be buffered up though.
				// TODO(yuzefovich): consider templating this.
				bufferedNull := bufferedGroup.firstTuple[colIdx].MaybeHasNulls() && bufferedGroup.firstTuple[colIdx].Nulls().NullAt(0)
				incomingNull := batch.ColVec(int(colIdx)).MaybeHasNulls() && batch.ColVec(int(colIdx)).Nulls().NullAt(tupleToLookAtIdx)
				if o.joinType.IsSetOpJoin() {
					if bufferedNull && incomingNull {
						// We have a NULL match, so move onto the next column.
						continue
					}
				}
				if bufferedNull || incomingNull {
					return true
				}
				bufferedCol := bufferedGroup.firstTuple[colIdx].Int32()
				prevVal := bufferedCol[0]
				col := batch.ColVec(int(colIdx)).Int32()
				curVal := col[tupleToLookAtIdx]
				var match bool

				{
					var cmpResult int

					{
						a, b := int64(prevVal), int64(curVal)
						if a < b {
							cmpResult = -1
						} else if a > b {
							cmpResult = 1
						} else {
							cmpResult = 0
						}
					}

					match = cmpResult == 0
				}

				if !match {
					return true
				}
				//
			case -1:
			default:
				// We perform this null check on every equality column of the first
				// buffered tuple regardless of the join type since it is done only once
				// per batch. In some cases (like INNER join, or LEFT OUTER join with the
				// right side being an input) this check will always return false since
				// nulls couldn't be buffered up though.
				// TODO(yuzefovich): consider templating this.
				bufferedNull := bufferedGroup.firstTuple[colIdx].MaybeHasNulls() && bufferedGroup.firstTuple[colIdx].Nulls().NullAt(0)
				incomingNull := batch.ColVec(int(colIdx)).MaybeHasNulls() && batch.ColVec(int(colIdx)).Nulls().NullAt(tupleToLookAtIdx)
				if o.joinType.IsSetOpJoin() {
					if bufferedNull && incomingNull {
						// We have a NULL match, so move onto the next column.
						continue
					}
				}
				if bufferedNull || incomingNull {
					return true
				}
				bufferedCol := bufferedGroup.firstTuple[colIdx].Int64()
				prevVal := bufferedCol[0]
				col := batch.ColVec(int(colIdx)).Int64()
				curVal := col[tupleToLookAtIdx]
				var match bool

				{
					var cmpResult int

					{
						a, b := int64(prevVal), int64(curVal)
						if a < b {
							cmpResult = -1
						} else if a > b {
							cmpResult = 1
						} else {
							cmpResult = 0
						}
					}

					match = cmpResult == 0
				}

				if !match {
					return true
				}
				//
			}
		//
		case types.FloatFamily:
			switch input.sourceTypes[colIdx].Width() {
			//
			case -1:
			default:
				// We perform this null check on every equality column of the first
				// buffered tuple regardless of the join type since it is done only once
				// per batch. In some cases (like INNER join, or LEFT OUTER join with the
				// right side being an input) this check will always return false since
				// nulls couldn't be buffered up though.
				// TODO(yuzefovich): consider templating this.
				bufferedNull := bufferedGroup.firstTuple[colIdx].MaybeHasNulls() && bufferedGroup.firstTuple[colIdx].Nulls().NullAt(0)
				incomingNull := batch.ColVec(int(colIdx)).MaybeHasNulls() && batch.ColVec(int(colIdx)).Nulls().NullAt(tupleToLookAtIdx)
				if o.joinType.IsSetOpJoin() {
					if bufferedNull && incomingNull {
						// We have a NULL match, so move onto the next column.
						continue
					}
				}
				if bufferedNull || incomingNull {
					return true
				}
				bufferedCol := bufferedGroup.firstTuple[colIdx].Float64()
				prevVal := bufferedCol[0]
				col := batch.ColVec(int(colIdx)).Float64()
				curVal := col[tupleToLookAtIdx]
				var match bool

				{
					var cmpResult int

					{
						a, b := float64(prevVal), float64(curVal)
						if a < b {
							cmpResult = -1
						} else if a > b {
							cmpResult = 1
						} else if a == b {
							cmpResult = 0
						} else if math.IsNaN(a) {
							if math.IsNaN(b) {
								cmpResult = 0
							} else {
								cmpResult = -1
							}
						} else {
							cmpResult = 1
						}
					}

					match = cmpResult == 0
				}

				if !match {
					return true
				}
				//
			}
		//
		case types.TimestampTZFamily:
			switch input.sourceTypes[colIdx].Width() {
			//
			case -1:
			default:
				// We perform this null check on every equality column of the first
				// buffered tuple regardless of the join type since it is done only once
				// per batch. In some cases (like INNER join, or LEFT OUTER join with the
				// right side being an input) this check will always return false since
				// nulls couldn't be buffered up though.
				// TODO(yuzefovich): consider templating this.
				bufferedNull := bufferedGroup.firstTuple[colIdx].MaybeHasNulls() && bufferedGroup.firstTuple[colIdx].Nulls().NullAt(0)
				incomingNull := batch.ColVec(int(colIdx)).MaybeHasNulls() && batch.ColVec(int(colIdx)).Nulls().NullAt(tupleToLookAtIdx)
				if o.joinType.IsSetOpJoin() {
					if bufferedNull && incomingNull {
						// We have a NULL match, so move onto the next column.
						continue
					}
				}
				if bufferedNull || incomingNull {
					return true
				}
				bufferedCol := bufferedGroup.firstTuple[colIdx].Timestamp()
				prevVal := bufferedCol[0]
				col := batch.ColVec(int(colIdx)).Timestamp()
				curVal := col[tupleToLookAtIdx]
				var match bool

				{
					var cmpResult int

					if prevVal.Before(curVal) {
						cmpResult = -1
					} else if curVal.Before(prevVal) {
						cmpResult = 1
					} else {
						cmpResult = 0
					}
					match = cmpResult == 0
				}

				if !match {
					return true
				}
				//
			}
		//
		case types.IntervalFamily:
			switch input.sourceTypes[colIdx].Width() {
			//
			case -1:
			default:
				// We perform this null check on every equality column of the first
				// buffered tuple regardless of the join type since it is done only once
				// per batch. In some cases (like INNER join, or LEFT OUTER join with the
				// right side being an input) this check will always return false since
				// nulls couldn't be buffered up though.
				// TODO(yuzefovich): consider templating this.
				bufferedNull := bufferedGroup.firstTuple[colIdx].MaybeHasNulls() && bufferedGroup.firstTuple[colIdx].Nulls().NullAt(0)
				incomingNull := batch.ColVec(int(colIdx)).MaybeHasNulls() && batch.ColVec(int(colIdx)).Nulls().NullAt(tupleToLookAtIdx)
				if o.joinType.IsSetOpJoin() {
					if bufferedNull && incomingNull {
						// We have a NULL match, so move onto the next column.
						continue
					}
				}
				if bufferedNull || incomingNull {
					return true
				}
				bufferedCol := bufferedGroup.firstTuple[colIdx].Interval()
				prevVal := bufferedCol[0]
				col := batch.ColVec(int(colIdx)).Interval()
				curVal := col[tupleToLookAtIdx]
				var match bool

				{
					var cmpResult int
					cmpResult = prevVal.Compare(curVal)
					match = cmpResult == 0
				}

				if !match {
					return true
				}
				//
			}
		//
		case typeconv.DatumVecCanonicalTypeFamily:
			switch input.sourceTypes[colIdx].Width() {
			//
			case -1:
			default:
				// We perform this null check on every equality column of the first
				// buffered tuple regardless of the join type since it is done only once
				// per batch. In some cases (like INNER join, or LEFT OUTER join with the
				// right side being an input) this check will always return false since
				// nulls couldn't be buffered up though.
				// TODO(yuzefovich): consider templating this.
				bufferedNull := bufferedGroup.firstTuple[colIdx].MaybeHasNulls() && bufferedGroup.firstTuple[colIdx].Nulls().NullAt(0)
				incomingNull := batch.ColVec(int(colIdx)).MaybeHasNulls() && batch.ColVec(int(colIdx)).Nulls().NullAt(tupleToLookAtIdx)
				if o.joinType.IsSetOpJoin() {
					if bufferedNull && incomingNull {
						// We have a NULL match, so move onto the next column.
						continue
					}
				}
				if bufferedNull || incomingNull {
					return true
				}
				bufferedCol := bufferedGroup.firstTuple[colIdx].Datum()
				prevVal := bufferedCol.Get(0)
				col := batch.ColVec(int(colIdx)).Datum()
				curVal := col.Get(tupleToLookAtIdx)
				var match bool

				{
					var cmpResult int

					cmpResult = prevVal.(*coldataext.Datum).CompareDatum(bufferedCol, curVal)

					match = cmpResult == 0
				}

				if !match {
					return true
				}
				//
			}
		//
		default:
			colexecerror.InternalError(fmt.Sprintf("unhandled type %s", input.sourceTypes[colIdx]))
		}
	}
	return false
}

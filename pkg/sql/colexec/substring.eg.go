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
	"context"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/col/coltypes"
	"github.com/cockroachdb/cockroach/pkg/col/coltypes/typeconv"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase/colexecerror"
	"github.com/cockroachdb/cockroach/pkg/sql/colmem"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/errors"
)

func newSubstringOperator(
	allocator *colmem.Allocator,
	typs []types.T,
	argumentCols []int,
	outputIdx int,
	input colexecbase.Operator,
) colexecbase.Operator {
	startType := &typs[argumentCols[1]]
	lengthType := &typs[argumentCols[2]]
	base := substringFunctionBase{
		OneInputNode: NewOneInputNode(input),
		allocator:    allocator,
		argumentCols: argumentCols,
		outputIdx:    outputIdx,
	}
	switch typeconv.FromColumnType(startType) {
	case coltypes.Int16:
		switch typeconv.FromColumnType(lengthType) {
		case coltypes.Int16:
			return &substringInt16Int16Operator{base}
		case coltypes.Int32:
			return &substringInt16Int32Operator{base}
		case coltypes.Int64:
			return &substringInt16Int64Operator{base}
		default:
			colexecerror.InternalError(errors.Errorf("unsupported length argument type %s", lengthType))
			// This code is unreachable, but the compiler cannot infer that.
			return nil
		}
	case coltypes.Int32:
		switch typeconv.FromColumnType(lengthType) {
		case coltypes.Int16:
			return &substringInt32Int16Operator{base}
		case coltypes.Int32:
			return &substringInt32Int32Operator{base}
		case coltypes.Int64:
			return &substringInt32Int64Operator{base}
		default:
			colexecerror.InternalError(errors.Errorf("unsupported length argument type %s", lengthType))
			// This code is unreachable, but the compiler cannot infer that.
			return nil
		}
	case coltypes.Int64:
		switch typeconv.FromColumnType(lengthType) {
		case coltypes.Int16:
			return &substringInt64Int16Operator{base}
		case coltypes.Int32:
			return &substringInt64Int32Operator{base}
		case coltypes.Int64:
			return &substringInt64Int64Operator{base}
		default:
			colexecerror.InternalError(errors.Errorf("unsupported length argument type %s", lengthType))
			// This code is unreachable, but the compiler cannot infer that.
			return nil
		}
	default:
		colexecerror.InternalError(errors.Errorf("unsupported start argument type %s", startType))
		// This code is unreachable, but the compiler cannot infer that.
		return nil
	}
}

type substringFunctionBase struct {
	OneInputNode
	allocator    *colmem.Allocator
	argumentCols []int
	outputIdx    int
}

func (s *substringFunctionBase) Init() {
	s.input.Init()
}

type substringInt16Int16Operator struct {
	substringFunctionBase
}

var _ colexecbase.Operator = &substringInt16Int16Operator{}

func (s *substringInt16Int16Operator) Next(ctx context.Context) coldata.Batch {
	batch := s.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}

	sel := batch.Selection()
	runeVec := batch.ColVec(s.argumentCols[0]).Bytes()
	startVec := batch.ColVec(s.argumentCols[1]).Int16()
	lengthVec := batch.ColVec(s.argumentCols[2]).Int16()
	outputVec := batch.ColVec(s.outputIdx)
	outputCol := outputVec.Bytes()
	s.allocator.PerformOperation(
		[]coldata.Vec{outputVec},
		func() {
			for i := 0; i < n; i++ {
				rowIdx := i
				if sel != nil {
					rowIdx = sel[i]
				}

				// The substring operator does not support nulls. If any of the arguments
				// are NULL, we output NULL.
				isNull := false
				for _, col := range s.argumentCols {
					if batch.ColVec(col).Nulls().NullAt(rowIdx) {
						isNull = true
						break
					}
				}
				if isNull {
					batch.ColVec(s.outputIdx).Nulls().SetNull(rowIdx)
					continue
				}

				runes := runeVec.Get(rowIdx)
				// Substring start is 1 indexed.
				start := int(startVec[rowIdx]) - 1
				length := int(lengthVec[rowIdx])
				if length < 0 {
					colexecerror.ExpectedError(errors.Errorf("negative substring length %d not allowed", length))
				}

				end := start + length
				// Check for integer overflow.
				if end < start {
					end = len(runes)
				} else if end < 0 {
					end = 0
				} else if end > len(runes) {
					end = len(runes)
				}

				if start < 0 {
					start = 0
				} else if start > len(runes) {
					start = len(runes)
				}
				outputCol.Set(rowIdx, runes[start:end])
			}
		},
	)
	// Although we didn't change the length of the batch, it is necessary to set
	// the length anyway (this helps maintaining the invariant of flat bytes).
	batch.SetLength(n)
	return batch
}

type substringInt16Int32Operator struct {
	substringFunctionBase
}

var _ colexecbase.Operator = &substringInt16Int32Operator{}

func (s *substringInt16Int32Operator) Next(ctx context.Context) coldata.Batch {
	batch := s.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}

	sel := batch.Selection()
	runeVec := batch.ColVec(s.argumentCols[0]).Bytes()
	startVec := batch.ColVec(s.argumentCols[1]).Int16()
	lengthVec := batch.ColVec(s.argumentCols[2]).Int32()
	outputVec := batch.ColVec(s.outputIdx)
	outputCol := outputVec.Bytes()
	s.allocator.PerformOperation(
		[]coldata.Vec{outputVec},
		func() {
			for i := 0; i < n; i++ {
				rowIdx := i
				if sel != nil {
					rowIdx = sel[i]
				}

				// The substring operator does not support nulls. If any of the arguments
				// are NULL, we output NULL.
				isNull := false
				for _, col := range s.argumentCols {
					if batch.ColVec(col).Nulls().NullAt(rowIdx) {
						isNull = true
						break
					}
				}
				if isNull {
					batch.ColVec(s.outputIdx).Nulls().SetNull(rowIdx)
					continue
				}

				runes := runeVec.Get(rowIdx)
				// Substring start is 1 indexed.
				start := int(startVec[rowIdx]) - 1
				length := int(lengthVec[rowIdx])
				if length < 0 {
					colexecerror.ExpectedError(errors.Errorf("negative substring length %d not allowed", length))
				}

				end := start + length
				// Check for integer overflow.
				if end < start {
					end = len(runes)
				} else if end < 0 {
					end = 0
				} else if end > len(runes) {
					end = len(runes)
				}

				if start < 0 {
					start = 0
				} else if start > len(runes) {
					start = len(runes)
				}
				outputCol.Set(rowIdx, runes[start:end])
			}
		},
	)
	// Although we didn't change the length of the batch, it is necessary to set
	// the length anyway (this helps maintaining the invariant of flat bytes).
	batch.SetLength(n)
	return batch
}

type substringInt16Int64Operator struct {
	substringFunctionBase
}

var _ colexecbase.Operator = &substringInt16Int64Operator{}

func (s *substringInt16Int64Operator) Next(ctx context.Context) coldata.Batch {
	batch := s.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}

	sel := batch.Selection()
	runeVec := batch.ColVec(s.argumentCols[0]).Bytes()
	startVec := batch.ColVec(s.argumentCols[1]).Int16()
	lengthVec := batch.ColVec(s.argumentCols[2]).Int64()
	outputVec := batch.ColVec(s.outputIdx)
	outputCol := outputVec.Bytes()
	s.allocator.PerformOperation(
		[]coldata.Vec{outputVec},
		func() {
			for i := 0; i < n; i++ {
				rowIdx := i
				if sel != nil {
					rowIdx = sel[i]
				}

				// The substring operator does not support nulls. If any of the arguments
				// are NULL, we output NULL.
				isNull := false
				for _, col := range s.argumentCols {
					if batch.ColVec(col).Nulls().NullAt(rowIdx) {
						isNull = true
						break
					}
				}
				if isNull {
					batch.ColVec(s.outputIdx).Nulls().SetNull(rowIdx)
					continue
				}

				runes := runeVec.Get(rowIdx)
				// Substring start is 1 indexed.
				start := int(startVec[rowIdx]) - 1
				length := int(lengthVec[rowIdx])
				if length < 0 {
					colexecerror.ExpectedError(errors.Errorf("negative substring length %d not allowed", length))
				}

				end := start + length
				// Check for integer overflow.
				if end < start {
					end = len(runes)
				} else if end < 0 {
					end = 0
				} else if end > len(runes) {
					end = len(runes)
				}

				if start < 0 {
					start = 0
				} else if start > len(runes) {
					start = len(runes)
				}
				outputCol.Set(rowIdx, runes[start:end])
			}
		},
	)
	// Although we didn't change the length of the batch, it is necessary to set
	// the length anyway (this helps maintaining the invariant of flat bytes).
	batch.SetLength(n)
	return batch
}

type substringInt32Int16Operator struct {
	substringFunctionBase
}

var _ colexecbase.Operator = &substringInt32Int16Operator{}

func (s *substringInt32Int16Operator) Next(ctx context.Context) coldata.Batch {
	batch := s.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}

	sel := batch.Selection()
	runeVec := batch.ColVec(s.argumentCols[0]).Bytes()
	startVec := batch.ColVec(s.argumentCols[1]).Int32()
	lengthVec := batch.ColVec(s.argumentCols[2]).Int16()
	outputVec := batch.ColVec(s.outputIdx)
	outputCol := outputVec.Bytes()
	s.allocator.PerformOperation(
		[]coldata.Vec{outputVec},
		func() {
			for i := 0; i < n; i++ {
				rowIdx := i
				if sel != nil {
					rowIdx = sel[i]
				}

				// The substring operator does not support nulls. If any of the arguments
				// are NULL, we output NULL.
				isNull := false
				for _, col := range s.argumentCols {
					if batch.ColVec(col).Nulls().NullAt(rowIdx) {
						isNull = true
						break
					}
				}
				if isNull {
					batch.ColVec(s.outputIdx).Nulls().SetNull(rowIdx)
					continue
				}

				runes := runeVec.Get(rowIdx)
				// Substring start is 1 indexed.
				start := int(startVec[rowIdx]) - 1
				length := int(lengthVec[rowIdx])
				if length < 0 {
					colexecerror.ExpectedError(errors.Errorf("negative substring length %d not allowed", length))
				}

				end := start + length
				// Check for integer overflow.
				if end < start {
					end = len(runes)
				} else if end < 0 {
					end = 0
				} else if end > len(runes) {
					end = len(runes)
				}

				if start < 0 {
					start = 0
				} else if start > len(runes) {
					start = len(runes)
				}
				outputCol.Set(rowIdx, runes[start:end])
			}
		},
	)
	// Although we didn't change the length of the batch, it is necessary to set
	// the length anyway (this helps maintaining the invariant of flat bytes).
	batch.SetLength(n)
	return batch
}

type substringInt32Int32Operator struct {
	substringFunctionBase
}

var _ colexecbase.Operator = &substringInt32Int32Operator{}

func (s *substringInt32Int32Operator) Next(ctx context.Context) coldata.Batch {
	batch := s.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}

	sel := batch.Selection()
	runeVec := batch.ColVec(s.argumentCols[0]).Bytes()
	startVec := batch.ColVec(s.argumentCols[1]).Int32()
	lengthVec := batch.ColVec(s.argumentCols[2]).Int32()
	outputVec := batch.ColVec(s.outputIdx)
	outputCol := outputVec.Bytes()
	s.allocator.PerformOperation(
		[]coldata.Vec{outputVec},
		func() {
			for i := 0; i < n; i++ {
				rowIdx := i
				if sel != nil {
					rowIdx = sel[i]
				}

				// The substring operator does not support nulls. If any of the arguments
				// are NULL, we output NULL.
				isNull := false
				for _, col := range s.argumentCols {
					if batch.ColVec(col).Nulls().NullAt(rowIdx) {
						isNull = true
						break
					}
				}
				if isNull {
					batch.ColVec(s.outputIdx).Nulls().SetNull(rowIdx)
					continue
				}

				runes := runeVec.Get(rowIdx)
				// Substring start is 1 indexed.
				start := int(startVec[rowIdx]) - 1
				length := int(lengthVec[rowIdx])
				if length < 0 {
					colexecerror.ExpectedError(errors.Errorf("negative substring length %d not allowed", length))
				}

				end := start + length
				// Check for integer overflow.
				if end < start {
					end = len(runes)
				} else if end < 0 {
					end = 0
				} else if end > len(runes) {
					end = len(runes)
				}

				if start < 0 {
					start = 0
				} else if start > len(runes) {
					start = len(runes)
				}
				outputCol.Set(rowIdx, runes[start:end])
			}
		},
	)
	// Although we didn't change the length of the batch, it is necessary to set
	// the length anyway (this helps maintaining the invariant of flat bytes).
	batch.SetLength(n)
	return batch
}

type substringInt32Int64Operator struct {
	substringFunctionBase
}

var _ colexecbase.Operator = &substringInt32Int64Operator{}

func (s *substringInt32Int64Operator) Next(ctx context.Context) coldata.Batch {
	batch := s.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}

	sel := batch.Selection()
	runeVec := batch.ColVec(s.argumentCols[0]).Bytes()
	startVec := batch.ColVec(s.argumentCols[1]).Int32()
	lengthVec := batch.ColVec(s.argumentCols[2]).Int64()
	outputVec := batch.ColVec(s.outputIdx)
	outputCol := outputVec.Bytes()
	s.allocator.PerformOperation(
		[]coldata.Vec{outputVec},
		func() {
			for i := 0; i < n; i++ {
				rowIdx := i
				if sel != nil {
					rowIdx = sel[i]
				}

				// The substring operator does not support nulls. If any of the arguments
				// are NULL, we output NULL.
				isNull := false
				for _, col := range s.argumentCols {
					if batch.ColVec(col).Nulls().NullAt(rowIdx) {
						isNull = true
						break
					}
				}
				if isNull {
					batch.ColVec(s.outputIdx).Nulls().SetNull(rowIdx)
					continue
				}

				runes := runeVec.Get(rowIdx)
				// Substring start is 1 indexed.
				start := int(startVec[rowIdx]) - 1
				length := int(lengthVec[rowIdx])
				if length < 0 {
					colexecerror.ExpectedError(errors.Errorf("negative substring length %d not allowed", length))
				}

				end := start + length
				// Check for integer overflow.
				if end < start {
					end = len(runes)
				} else if end < 0 {
					end = 0
				} else if end > len(runes) {
					end = len(runes)
				}

				if start < 0 {
					start = 0
				} else if start > len(runes) {
					start = len(runes)
				}
				outputCol.Set(rowIdx, runes[start:end])
			}
		},
	)
	// Although we didn't change the length of the batch, it is necessary to set
	// the length anyway (this helps maintaining the invariant of flat bytes).
	batch.SetLength(n)
	return batch
}

type substringInt64Int16Operator struct {
	substringFunctionBase
}

var _ colexecbase.Operator = &substringInt64Int16Operator{}

func (s *substringInt64Int16Operator) Next(ctx context.Context) coldata.Batch {
	batch := s.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}

	sel := batch.Selection()
	runeVec := batch.ColVec(s.argumentCols[0]).Bytes()
	startVec := batch.ColVec(s.argumentCols[1]).Int64()
	lengthVec := batch.ColVec(s.argumentCols[2]).Int16()
	outputVec := batch.ColVec(s.outputIdx)
	outputCol := outputVec.Bytes()
	s.allocator.PerformOperation(
		[]coldata.Vec{outputVec},
		func() {
			for i := 0; i < n; i++ {
				rowIdx := i
				if sel != nil {
					rowIdx = sel[i]
				}

				// The substring operator does not support nulls. If any of the arguments
				// are NULL, we output NULL.
				isNull := false
				for _, col := range s.argumentCols {
					if batch.ColVec(col).Nulls().NullAt(rowIdx) {
						isNull = true
						break
					}
				}
				if isNull {
					batch.ColVec(s.outputIdx).Nulls().SetNull(rowIdx)
					continue
				}

				runes := runeVec.Get(rowIdx)
				// Substring start is 1 indexed.
				start := int(startVec[rowIdx]) - 1
				length := int(lengthVec[rowIdx])
				if length < 0 {
					colexecerror.ExpectedError(errors.Errorf("negative substring length %d not allowed", length))
				}

				end := start + length
				// Check for integer overflow.
				if end < start {
					end = len(runes)
				} else if end < 0 {
					end = 0
				} else if end > len(runes) {
					end = len(runes)
				}

				if start < 0 {
					start = 0
				} else if start > len(runes) {
					start = len(runes)
				}
				outputCol.Set(rowIdx, runes[start:end])
			}
		},
	)
	// Although we didn't change the length of the batch, it is necessary to set
	// the length anyway (this helps maintaining the invariant of flat bytes).
	batch.SetLength(n)
	return batch
}

type substringInt64Int32Operator struct {
	substringFunctionBase
}

var _ colexecbase.Operator = &substringInt64Int32Operator{}

func (s *substringInt64Int32Operator) Next(ctx context.Context) coldata.Batch {
	batch := s.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}

	sel := batch.Selection()
	runeVec := batch.ColVec(s.argumentCols[0]).Bytes()
	startVec := batch.ColVec(s.argumentCols[1]).Int64()
	lengthVec := batch.ColVec(s.argumentCols[2]).Int32()
	outputVec := batch.ColVec(s.outputIdx)
	outputCol := outputVec.Bytes()
	s.allocator.PerformOperation(
		[]coldata.Vec{outputVec},
		func() {
			for i := 0; i < n; i++ {
				rowIdx := i
				if sel != nil {
					rowIdx = sel[i]
				}

				// The substring operator does not support nulls. If any of the arguments
				// are NULL, we output NULL.
				isNull := false
				for _, col := range s.argumentCols {
					if batch.ColVec(col).Nulls().NullAt(rowIdx) {
						isNull = true
						break
					}
				}
				if isNull {
					batch.ColVec(s.outputIdx).Nulls().SetNull(rowIdx)
					continue
				}

				runes := runeVec.Get(rowIdx)
				// Substring start is 1 indexed.
				start := int(startVec[rowIdx]) - 1
				length := int(lengthVec[rowIdx])
				if length < 0 {
					colexecerror.ExpectedError(errors.Errorf("negative substring length %d not allowed", length))
				}

				end := start + length
				// Check for integer overflow.
				if end < start {
					end = len(runes)
				} else if end < 0 {
					end = 0
				} else if end > len(runes) {
					end = len(runes)
				}

				if start < 0 {
					start = 0
				} else if start > len(runes) {
					start = len(runes)
				}
				outputCol.Set(rowIdx, runes[start:end])
			}
		},
	)
	// Although we didn't change the length of the batch, it is necessary to set
	// the length anyway (this helps maintaining the invariant of flat bytes).
	batch.SetLength(n)
	return batch
}

type substringInt64Int64Operator struct {
	substringFunctionBase
}

var _ colexecbase.Operator = &substringInt64Int64Operator{}

func (s *substringInt64Int64Operator) Next(ctx context.Context) coldata.Batch {
	batch := s.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}

	sel := batch.Selection()
	runeVec := batch.ColVec(s.argumentCols[0]).Bytes()
	startVec := batch.ColVec(s.argumentCols[1]).Int64()
	lengthVec := batch.ColVec(s.argumentCols[2]).Int64()
	outputVec := batch.ColVec(s.outputIdx)
	outputCol := outputVec.Bytes()
	s.allocator.PerformOperation(
		[]coldata.Vec{outputVec},
		func() {
			for i := 0; i < n; i++ {
				rowIdx := i
				if sel != nil {
					rowIdx = sel[i]
				}

				// The substring operator does not support nulls. If any of the arguments
				// are NULL, we output NULL.
				isNull := false
				for _, col := range s.argumentCols {
					if batch.ColVec(col).Nulls().NullAt(rowIdx) {
						isNull = true
						break
					}
				}
				if isNull {
					batch.ColVec(s.outputIdx).Nulls().SetNull(rowIdx)
					continue
				}

				runes := runeVec.Get(rowIdx)
				// Substring start is 1 indexed.
				start := int(startVec[rowIdx]) - 1
				length := int(lengthVec[rowIdx])
				if length < 0 {
					colexecerror.ExpectedError(errors.Errorf("negative substring length %d not allowed", length))
				}

				end := start + length
				// Check for integer overflow.
				if end < start {
					end = len(runes)
				} else if end < 0 {
					end = 0
				} else if end > len(runes) {
					end = len(runes)
				}

				if start < 0 {
					start = 0
				} else if start > len(runes) {
					start = len(runes)
				}
				outputCol.Set(rowIdx, runes[start:end])
			}
		},
	)
	// Although we didn't change the length of the batch, it is necessary to set
	// the length anyway (this helps maintaining the invariant of flat bytes).
	batch.SetLength(n)
	return batch
}

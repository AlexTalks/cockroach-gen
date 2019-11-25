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
	"fmt"
	"time"

	"github.com/cockroachdb/apd"
	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/colexec/execerror"
	"github.com/cockroachdb/cockroach/pkg/sql/colexec/execgen"
	"github.com/cockroachdb/cockroach/pkg/sql/colexec/typeconv"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
)

// Use execgen package to remove unused import warning.
var _ interface{} = execgen.UNSAFEGET

// EncDatumRowsToColVec converts one column from EncDatumRows to a column
// vector. columnIdx is the 0-based index of the column in the EncDatumRows.
func EncDatumRowsToColVec(
	allocator *Allocator,
	rows sqlbase.EncDatumRows,
	vec coldata.Vec,
	columnIdx int,
	columnType *types.T,
	alloc *sqlbase.DatumAlloc,
) error {
	var err error
	switch columnType.Family() {
	case types.BytesFamily:
		allocator.performOperation(
			[]coldata.Vec{vec},
			func() {
				col := vec.Bytes()
				datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
				for i := range rows {
					row := rows[i]
					if row[columnIdx].Datum == nil {
						if err = row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
							break
						}
					}
					datum := row[columnIdx].Datum
					if datum == tree.DNull {
						vec.Nulls().SetNull(uint16(i))
					} else {
						v, err := datumToPhysicalFn(datum)
						if err != nil {
							break
						}

						castV := v.([]byte)
						col.Set(i, castV)
					}
				}
			},
		)
		return err
	case types.BoolFamily:
		allocator.performOperation(
			[]coldata.Vec{vec},
			func() {
				col := vec.Bool()
				datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
				for i := range rows {
					row := rows[i]
					if row[columnIdx].Datum == nil {
						if err = row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
							break
						}
					}
					datum := row[columnIdx].Datum
					if datum == tree.DNull {
						vec.Nulls().SetNull(uint16(i))
					} else {
						v, err := datumToPhysicalFn(datum)
						if err != nil {
							break
						}

						castV := v.(bool)
						col[i] = castV
					}
				}
			},
		)
		return err
	case types.DateFamily:
		allocator.performOperation(
			[]coldata.Vec{vec},
			func() {
				col := vec.Int64()
				datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
				for i := range rows {
					row := rows[i]
					if row[columnIdx].Datum == nil {
						if err = row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
							break
						}
					}
					datum := row[columnIdx].Datum
					if datum == tree.DNull {
						vec.Nulls().SetNull(uint16(i))
					} else {
						v, err := datumToPhysicalFn(datum)
						if err != nil {
							break
						}

						castV := v.(int64)
						col[i] = castV
					}
				}
			},
		)
		return err
	case types.StringFamily:
		allocator.performOperation(
			[]coldata.Vec{vec},
			func() {
				col := vec.Bytes()
				datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
				for i := range rows {
					row := rows[i]
					if row[columnIdx].Datum == nil {
						if err = row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
							break
						}
					}
					datum := row[columnIdx].Datum
					if datum == tree.DNull {
						vec.Nulls().SetNull(uint16(i))
					} else {
						v, err := datumToPhysicalFn(datum)
						if err != nil {
							break
						}

						castV := v.([]byte)
						col.Set(i, castV)
					}
				}
			},
		)
		return err
	case types.OidFamily:
		allocator.performOperation(
			[]coldata.Vec{vec},
			func() {
				col := vec.Int64()
				datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
				for i := range rows {
					row := rows[i]
					if row[columnIdx].Datum == nil {
						if err = row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
							break
						}
					}
					datum := row[columnIdx].Datum
					if datum == tree.DNull {
						vec.Nulls().SetNull(uint16(i))
					} else {
						v, err := datumToPhysicalFn(datum)
						if err != nil {
							break
						}

						castV := v.(int64)
						col[i] = castV
					}
				}
			},
		)
		return err
	case types.TimestampFamily:
		allocator.performOperation(
			[]coldata.Vec{vec},
			func() {
				col := vec.Timestamp()
				datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
				for i := range rows {
					row := rows[i]
					if row[columnIdx].Datum == nil {
						if err = row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
							break
						}
					}
					datum := row[columnIdx].Datum
					if datum == tree.DNull {
						vec.Nulls().SetNull(uint16(i))
					} else {
						v, err := datumToPhysicalFn(datum)
						if err != nil {
							break
						}

						castV := v.(time.Time)
						col[i] = castV
					}
				}
			},
		)
		return err
	case types.UuidFamily:
		allocator.performOperation(
			[]coldata.Vec{vec},
			func() {
				col := vec.Bytes()
				datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
				for i := range rows {
					row := rows[i]
					if row[columnIdx].Datum == nil {
						if err = row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
							break
						}
					}
					datum := row[columnIdx].Datum
					if datum == tree.DNull {
						vec.Nulls().SetNull(uint16(i))
					} else {
						v, err := datumToPhysicalFn(datum)
						if err != nil {
							break
						}

						castV := v.([]byte)
						col.Set(i, castV)
					}
				}
			},
		)
		return err
	case types.FloatFamily:
		switch columnType.Width() {
		case 64:
			allocator.performOperation(
				[]coldata.Vec{vec},
				func() {
					col := vec.Float64()
					datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
					for i := range rows {
						row := rows[i]
						if row[columnIdx].Datum == nil {
							if err = row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
								break
							}
						}
						datum := row[columnIdx].Datum
						if datum == tree.DNull {
							vec.Nulls().SetNull(uint16(i))
						} else {
							v, err := datumToPhysicalFn(datum)
							if err != nil {
								break
							}

							castV := v.(float64)
							col[i] = castV
						}
					}
				},
			)
			return err
		case 32:
			allocator.performOperation(
				[]coldata.Vec{vec},
				func() {
					col := vec.Float64()
					datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
					for i := range rows {
						row := rows[i]
						if row[columnIdx].Datum == nil {
							if err = row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
								break
							}
						}
						datum := row[columnIdx].Datum
						if datum == tree.DNull {
							vec.Nulls().SetNull(uint16(i))
						} else {
							v, err := datumToPhysicalFn(datum)
							if err != nil {
								break
							}

							castV := v.(float64)
							col[i] = castV
						}
					}
				},
			)
			return err
		default:
			execerror.VectorizedInternalPanic(fmt.Sprintf("unsupported width %d for column type %s", columnType.Width(), columnType.String()))
		}
	case types.IntFamily:
		switch columnType.Width() {
		case 16:
			allocator.performOperation(
				[]coldata.Vec{vec},
				func() {
					col := vec.Int16()
					datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
					for i := range rows {
						row := rows[i]
						if row[columnIdx].Datum == nil {
							if err = row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
								break
							}
						}
						datum := row[columnIdx].Datum
						if datum == tree.DNull {
							vec.Nulls().SetNull(uint16(i))
						} else {
							v, err := datumToPhysicalFn(datum)
							if err != nil {
								break
							}

							castV := v.(int16)
							col[i] = castV
						}
					}
				},
			)
			return err
		case 32:
			allocator.performOperation(
				[]coldata.Vec{vec},
				func() {
					col := vec.Int32()
					datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
					for i := range rows {
						row := rows[i]
						if row[columnIdx].Datum == nil {
							if err = row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
								break
							}
						}
						datum := row[columnIdx].Datum
						if datum == tree.DNull {
							vec.Nulls().SetNull(uint16(i))
						} else {
							v, err := datumToPhysicalFn(datum)
							if err != nil {
								break
							}

							castV := v.(int32)
							col[i] = castV
						}
					}
				},
			)
			return err
		case 64:
			allocator.performOperation(
				[]coldata.Vec{vec},
				func() {
					col := vec.Int64()
					datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
					for i := range rows {
						row := rows[i]
						if row[columnIdx].Datum == nil {
							if err = row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
								break
							}
						}
						datum := row[columnIdx].Datum
						if datum == tree.DNull {
							vec.Nulls().SetNull(uint16(i))
						} else {
							v, err := datumToPhysicalFn(datum)
							if err != nil {
								break
							}

							castV := v.(int64)
							col[i] = castV
						}
					}
				},
			)
			return err
		default:
			execerror.VectorizedInternalPanic(fmt.Sprintf("unsupported width %d for column type %s", columnType.Width(), columnType.String()))
		}
	case types.DecimalFamily:
		allocator.performOperation(
			[]coldata.Vec{vec},
			func() {
				col := vec.Decimal()
				datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
				for i := range rows {
					row := rows[i]
					if row[columnIdx].Datum == nil {
						if err = row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
							break
						}
					}
					datum := row[columnIdx].Datum
					if datum == tree.DNull {
						vec.Nulls().SetNull(uint16(i))
					} else {
						v, err := datumToPhysicalFn(datum)
						if err != nil {
							break
						}

						castV := v.(apd.Decimal)
						col[i].Set(&castV)
					}
				}
			},
		)
		return err
	default:
		execerror.VectorizedInternalPanic(fmt.Sprintf("unsupported column type %s", columnType.String()))
	}
	return nil
}

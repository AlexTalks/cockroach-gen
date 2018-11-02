// Code generated by execgen; DO NOT EDIT.

package exec

import (
	"fmt"

	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
	"github.com/cockroachdb/cockroach/pkg/util/encoding"
)

// EncDatumRowsToColVec converts one column from EncDatumRows to a column
// vector. columnIdx is the 0-based index of the column in the EncDatumRows.
func EncDatumRowsToColVec(
	rows sqlbase.EncDatumRows,
	vec ColVec,
	columnIdx int,
	columnType *sqlbase.ColumnType,
	alloc *sqlbase.DatumAlloc,
) error {
	nRows := uint16(len(rows))
	// TODO(solon): Make this chain of conditionals more efficient: either a
	// switch statement or even better a lookup table on SemanticType. Also get
	// rid of the somewhat dubious assumption that Width is unset (0) for column
	// types where it does not apply.

	if columnType.SemanticType == sqlbase.ColumnType_BOOL && columnType.Width == 0 {
		col := vec.Bool()
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				col[i] = bool(*datum.(*tree.DBool))
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_STRING && columnType.Width == 0 {
		col := vec.Bytes()
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				col[i] = encoding.UnsafeConvertStringToBytes(string(*datum.(*tree.DString)))
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_NAME && columnType.Width == 0 {
		col := vec.Bytes()
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				col[i] = encoding.UnsafeConvertStringToBytes(string(*datum.(*tree.DString)))
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_OID && columnType.Width == 0 {
		col := vec.Int64()
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				col[i] = int64(datum.(*tree.DOid).DInt)
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_DECIMAL && columnType.Width == 0 {
		col := vec.Decimal()
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				col[i] = datum.(*tree.DDecimal).Decimal
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_INT && columnType.Width == 0 {
		col := vec.Int64()
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				col[i] = int64(*datum.(*tree.DInt))
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_INT && columnType.Width == 8 {
		col := vec.Int8()
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				col[i] = int8(*datum.(*tree.DInt))
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_INT && columnType.Width == 16 {
		col := vec.Int16()
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				col[i] = int16(*datum.(*tree.DInt))
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_INT && columnType.Width == 32 {
		col := vec.Int32()
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				col[i] = int32(*datum.(*tree.DInt))
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_INT && columnType.Width == 64 {
		col := vec.Int64()
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				col[i] = int64(*datum.(*tree.DInt))
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_FLOAT && columnType.Width == 0 {
		col := vec.Float64()
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				col[i] = float64(*datum.(*tree.DFloat))
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_DATE && columnType.Width == 0 {
		col := vec.Int64()
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				col[i] = int64(*datum.(*tree.DDate))
			}
		}
		return nil
	}

	if columnType.SemanticType == sqlbase.ColumnType_BYTES && columnType.Width == 0 {
		col := vec.Bytes()
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				col[i] = encoding.UnsafeConvertStringToBytes(string(*datum.(*tree.DBytes)))
			}
		}
		return nil
	}

	panic(fmt.Sprintf("Unsupported column type and width: %s, %d", columnType.SQLString(), columnType.Width))
}

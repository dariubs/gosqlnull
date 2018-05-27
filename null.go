// Package null : better nullable types for golang sql for json outputs
package null

import (
	"database/sql"
	"encoding/json"
)

// Bool type for Nullable bool
type Bool struct {
	sql.NullBool
}

// String type for Nullable string
type String struct {
	sql.NullString
}

// Float64 type for Nullable float64
type Float64 struct {
	sql.NullFloat64
}

// Int64 type for Nullable int64
type Int64 struct {
	sql.NullInt64
}

// MarshalJSON for Bool data
// Flat output for Marshal JSON for Nullable bool
func (out Bool) MarshalJSON() ([]byte, error) {
	if out.Valid {
		return json.Marshal(out.Bool)
	}
	return json.Marshal(nil)
}

// MarshalJSON for String data
// Flat output for Marshal JSON for Nullable string
func (out String) MarshalJSON() ([]byte, error) {
	if out.Valid {
		return json.Marshal(out.String)
	}
	return json.Marshal(nil)
}

// MarshalJSON for Float64 data
// Flat output for Marshal JSON for Nullable float64
func (out Float64) MarshalJSON() ([]byte, error) {
	if out.Valid {
		return json.Marshal(out.Float64)
	}
	return json.Marshal(nil)
}

// MarshalJSON for Int64 data
// Flat output for Marshal JSON for Nullable int64
func (out Int64) MarshalJSON() ([]byte, error) {
	if out.Valid {
		return json.Marshal(out.Int64)
	}
	return json.Marshal(nil)
}

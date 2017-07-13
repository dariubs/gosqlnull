// better nullable types for golang sql for json outputs
package null

import(
  "database/sql"
  "encoding/json"
)

type Bool struct {
  sql.NullBool
}

type String struct {
  sql.NullString
}

type Float64 struct {
  sql.NullFloat64
}

type Int64 struct {
  sql.NullInt64
}

func (out Bool) MarshalJSON() ([]byte, error) {
  if out.Valid {
    return json.Marshal(out.Bool)
  }
  return json.Marshal(nil)
}

func (out String) MarshalJSON() ([]byte, error) {
  if out.Valid {
    return json.Marshal(out.String)
  }
  return json.Marshal(nil)
}

func (out Float64) MarshalJSON() ([]byte, error) {
  if out.Valid {
    return json.Marshal(out.Float64)
  }
  return json.Marshal(nil)
}

func (out Int64) MarshalJSON() ([]byte, error) {
  if out.Valid {
    return json.Marshal(out.Int64)
  }
  return json.Marshal(nil)
}

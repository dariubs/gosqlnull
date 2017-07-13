// better nullable types for golang sql for json outputs
package null

import(
  "database/sql"
  "encoding/json"
)

type String struct {
  sql.NullString
}

type Int64 struct {
  sql.NullInt64
}

func (out String) MarshalJSON() ([]byte, error) {
  if out.Valid {
    return json.Marshal(out.String)
  }
  return json.Marshal(nil)
}

func (out Int64) MarshalJSON() ([]byte, error) {
  if out.Valid {
    return json.Marshal(out.Int64)
  }
  return json.Marshal(nil)
}

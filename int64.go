package nulltype

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type NullInt64 struct {
	i sql.NullInt64
}

func (i *NullInt64) Valid() bool {
	return i.i.Valid
}

func (i *NullInt64) Value() int64 {
	return i.i.Int64
}

func (i *NullInt64) Reset() {
	i.i.Int64 = 0
	i.i.Valid = false
}

func (i *NullInt64) Set(value int64) {
	i.i.Valid = true
	i.i.Int64 = value
}

func (i *NullInt64) Scan(value interface{}) error {
	return i.i.Scan(value)
}

func (i NullInt64) String() string {
	if !i.i.Valid {
		return "nil"
	}
	return fmt.Sprint(i.i.Int64)
}

func (i NullInt64) MarshalJSON() ([]byte, error) {
	if !i.i.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(i.i.Int64)
}

func (i *NullInt64) UnmarshalJSON(data []byte) error {
	var value *int64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	i.i.Valid = value != nil
	if value == nil {
		i.i.Int64 = 0
	} else {
		i.i.Int64 = *value
	}
	return nil
}

package nulltype

import (
	"database/sql"
	"encoding/json"
)

type NullBool struct {
	b sql.NullBool
}

func (b *NullBool) Valid() bool {
	return b.b.Valid
}

func (b *NullBool) Value() bool {
	return b.b.Bool
}

func (b *NullBool) Reset() {
	b.b.Bool = false
	b.b.Valid = false
}

func (b *NullBool) Set(value bool) {
	b.b.Valid = true
	b.b.Bool = value
}

func (b *NullBool) Scan(value interface{}) error {
	return b.b.Scan(value)
}

func (b NullBool) String() string {
	if !b.b.Valid {
		return "nil"
	}
	if b.b.Bool {
		return "true"
	}
	return "false"
}

func (b NullBool) MarshalJSON() ([]byte, error) {
	if !b.b.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(b.b.Bool)
}

func (b *NullBool) UnmarshalJSON(data []byte) error {
	var value *bool
	if err := json.Unmarshal(data, &value); err != nil {
		b.b.Bool = false
		b.b.Valid = false
		return err
	}
	b.b.Valid = value != nil
	if value == nil {
		b.b.Bool = false
	} else {
		b.b.Bool = true
	}
	return nil
}

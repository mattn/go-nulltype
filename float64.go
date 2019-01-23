package nulltype

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type NullFloat64 struct {
	f sql.NullFloat64
}

func (f *NullFloat64) Valid() bool {
	return f.f.Valid
}

func (f *NullFloat64) Value() float64 {
	return f.f.Float64
}

func (f *NullFloat64) Reset() {
	f.f.Float64 = 0
	f.f.Valid = false
}

func (f *NullFloat64) Set(value float64) {
	f.f.Valid = true
	f.f.Float64 = value
}

func (f *NullFloat64) Scan(value interface{}) error {
	return f.f.Scan(value)
}

func (f NullFloat64) String() string {
	if !f.f.Valid {
		return "nil"
	}
	return fmt.Sprint(f.f.Float64)
}

func (f NullFloat64) MarshalJSON() ([]byte, error) {
	if !f.f.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(f.f.Float64)
}

func (f *NullFloat64) UnmarshalJSON(data []byte) error {
	var value *float64
	if err := json.Unmarshal(data, &value); err != nil {
		f.f.Float64 = 0
		f.f.Valid = false
		return err
	}
	f.f.Valid = value != nil
	if value == nil {
		f.f.Float64 = 0
	} else {
		f.f.Float64 = *value
	}
	return nil
}

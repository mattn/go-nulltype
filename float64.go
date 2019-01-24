package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// NullFloat64 is null friendly type for float64.
type NullFloat64 struct {
	f sql.NullFloat64
}

// NullFloat64Of return NullFloat64 that he value is set.
func NullFloat64Of(value float64) NullFloat64 {
	var s NullFloat64
	s.Set(value)
	return s
}

// Valid return the value is valid. If true, it is not null value.
func (f *NullFloat64) Valid() bool {
	return f.f.Valid
}

// Float64Value return the value.
func (f *NullFloat64) Float64Value() float64 {
	return f.f.Float64
}

// Reset set nil to the value.
func (f *NullFloat64) Reset() {
	f.f.Float64 = 0
	f.f.Valid = false
}

// Set set the value.
func (f *NullFloat64) Set(value float64) *NullFloat64 {
	f.f.Valid = true
	f.f.Float64 = value
	return f
}

// Scan is a method for database/sql.
func (f *NullFloat64) Scan(value interface{}) error {
	return f.f.Scan(value)
}

// String return string indicated the value.
func (f NullFloat64) String() string {
	if !f.f.Valid {
		return "nil"
	}
	return fmt.Sprint(f.f.Float64)
}

// MarshalJSON encode the value to JSON.
func (f NullFloat64) MarshalJSON() ([]byte, error) {
	if !f.f.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(f.f.Float64)
}

// UnmarshalJSON decode data to the value.
func (f *NullFloat64) UnmarshalJSON(data []byte) error {
	var value *float64
	if err := json.Unmarshal(data, &value); err != nil {
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

// Value implement driver.Valuer.
func (f NullFloat64) Value() (driver.Value, error) {
	return f.f.Float64, nil
}

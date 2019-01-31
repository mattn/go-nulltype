package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// NullBool is null friendly type for bool.
type NullBool struct {
	b sql.NullBool
}

// NullBoolOf return NullBool that he value is set.
func NullBoolOf(value bool) NullBool {
	var b NullBool
	b.Set(value)
	return b
}

// Valid return the value is valid. If true, it is not null value.
func (b *NullBool) Valid() bool {
	return b.b.Valid
}

// BoolValue return the value.
func (b *NullBool) BoolValue() bool {
	return b.b.Bool
}

// Reset set nil to the value.
func (b *NullBool) Reset() {
	b.b.Bool = false
	b.b.Valid = false
}

// Set set the value.
func (b *NullBool) Set(value bool) *NullBool {
	b.b.Valid = true
	b.b.Bool = value
	return b
}

// Scan is a method for database/sql.
func (b *NullBool) Scan(value interface{}) error {
	return b.b.Scan(value)
}

// String return string indicated the value.
func (b NullBool) String() string {
	if !b.b.Valid {
		return ""
	}
	if b.b.Bool {
		return "true"
	}
	return "false"
}

// MarshalJSON encode the value to JSON.
func (b NullBool) MarshalJSON() ([]byte, error) {
	if !b.b.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(b.b.Bool)
}

// UnmarshalJSON decode data to the value.
func (b *NullBool) UnmarshalJSON(data []byte) error {
	var value *bool
	if err := json.Unmarshal(data, &value); err != nil {
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

// Value implement driver.Valuer.
func (b NullBool) Value() (driver.Value, error) {
	if !b.Valid() {
		return nil, nil
	}
	return b.b.Bool, nil
}

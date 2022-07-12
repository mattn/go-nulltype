package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// NullString is null friendly type for string.
type NullString struct {
	s sql.NullString
}

// NullStringOf return NullString that he value is set.
func NullStringOf(value string) NullString {
	var s NullString
	s.Set(value)
	return s
}

// Valid return the value is valid. If true, it is not null value.
func (s NullString) Valid() bool {
	return s.s.Valid
}

// StringValue return the value.
func (s *NullString) StringValue() string {
	return s.s.String
}

// Reset set nil to the value.
func (s *NullString) Reset() {
	s.s.String = ""
	s.s.Valid = false
}

// Set set the value.
func (s *NullString) Set(value string) {
	s.s.Valid = true
	s.s.String = value
}

// Scan is a method for database/sql.
func (s *NullString) Scan(value interface{}) error {
	return s.s.Scan(value)
}

// String return string indicated the value.
func (s NullString) String() string {
	if !s.s.Valid {
		return ""
	}
	return s.s.String
}

// MarshalJSON encode the value to JSON.
func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.s.String)
}

// UnmarshalJSON decode data to the value.
func (s *NullString) UnmarshalJSON(data []byte) error {
	var value *string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	s.s.Valid = value != nil
	if value == nil {
		s.s.String = ""
	} else {
		s.s.String = *value
	}
	return nil
}

// Value implement driver.Valuer.
func (s NullString) Value() (driver.Value, error) {
	if !s.Valid() {
		return nil, nil
	}
	return s.s.String, nil
}

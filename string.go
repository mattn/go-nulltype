package nulltype

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	s sql.NullString
}

func (s *NullString) Valid() bool {
	return s.s.Valid
}

func (s *NullString) Value() string {
	return s.s.String
}

func (s *NullString) Reset() {
	s.s.String = ""
	s.s.Valid = false
}

func (s *NullString) Set(value string) {
	s.s.Valid = true
	s.s.String = value
}

func (s *NullString) Scan(value interface{}) error {
	return s.s.Scan(value)
}

func (s NullString) String() string {
	if !s.s.Valid {
		return "nil"
	}
	return s.s.String
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.s.String)
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	var value *string
	if err := json.Unmarshal(data, &value); err != nil {
		s.s.String = ""
		s.s.Valid = false
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

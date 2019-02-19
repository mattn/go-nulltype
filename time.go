package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"time"
)

// NullTime is null friendly type for string.
type NullTime struct {
	t time.Time
	v bool // Valid is true if Time is not NULL
}

// NullTimeOf return NullTime that he value is set.
func NullTimeOf(value time.Time) NullTime {
	var t NullTime
	t.Set(value)
	return t
}

// Valid return the value is valid. If true, it is not null value.
func (t *NullTime) Valid() bool {
	return t.v
}

// TimeValue return the value.
func (t *NullTime) TimeValue() time.Time {
	return t.t
}

// Reset set nil to the value.
func (t *NullTime) Reset() {
	t.t = time.Unix(0, 0)
	t.v = false
}

// Set set the value.
func (t *NullTime) Set(value time.Time) {
	t.v = true
	t.t = value
}

var timestampFormats = []string{
	"2006-01-02 15:04:05.999999999-07:00",
	"2006-01-02T15:04:05.999999999-07:00",
	"2006-01-02 15:04:05.999999999",
	"2006-01-02T15:04:05.999999999",
	"2006-01-02 15:04:05",
	"2006-01-02T15:04:05",
	"2006-01-02 15:04",
	"2006-01-02T15:04",
	"2006-01-02",
	"2006/01/02 15:04:05",
}

// Scan is a method for database/sql.
func (t *NullTime) Scan(value interface{}) error {
	t.t, t.v = value.(time.Time)
	if t.v {
		return nil
	}
	var ns sql.NullString
	if err := ns.Scan(value); err != nil {
		return err
	}
	for _, tf := range timestampFormats {
		if tt, err := time.Parse(tf, ns.String); err == nil {
			t.t = tt
			t.v = true
			return nil
		}

	}
	return nil
}

// Time return string indicated the value.
func (t NullTime) String() string {
	if !t.v {
		return ""
	}
	return t.t.Format("2006/01/02 15:04:05")
}

// MarshalJSON encode the value to JSON.
func (t NullTime) MarshalJSON() ([]byte, error) {
	if !t.v {
		return []byte("null"), nil
	}
	return json.Marshal(t.t.Format(time.RFC3339))
}

// UnmarshalJSON decode data to the value.
func (t *NullTime) UnmarshalJSON(data []byte) error {
	var value *string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	t.v = value != nil
	if value == nil {
		t.t = time.Unix(0, 0)
	} else {
		tt, err := time.Parse(time.RFC3339, *value)
		if err != nil {
			return err
		}
		t.t = tt
	}
	return nil
}

// Value implement driver.Valuer.
func (t NullTime) Value() (driver.Value, error) {
	if !t.Valid() {
		return nil, nil
	}
	return t.t, nil
}

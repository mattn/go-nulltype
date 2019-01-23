package nulltype

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestNullStringStringer(t *testing.T) {
	var s NullString

	want := "nil"
	got := fmt.Sprint(s)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = "foo"
	s.Set("foo")
	got = fmt.Sprint(s)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = "bar"
	s = NullStringOf("bar")
	got = fmt.Sprint(s)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = "nil"
	s.Reset()
	got = fmt.Sprint(s)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullStringMarshalJSON(t *testing.T) {
	var s NullString

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(s)
	if err != nil {
		t.Fatal(err)
	}

	want := "null"
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	buf.Reset()

	s.Set("foo")
	err = json.NewEncoder(&buf).Encode(s)
	if err != nil {
		t.Fatal(err)
	}

	want = `"foo"`
	got = strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullStringUnmarshalJSON(t *testing.T) {
	var s NullString

	err := json.NewDecoder(strings.NewReader("null")).Decode(&s)
	if err != nil {
		t.Fatal(err)
	}

	if s.Valid() {
		t.Fatalf("must be null but got %v", s)
	}

	err = json.NewDecoder(strings.NewReader(`"foo"`)).Decode(&s)
	if err != nil {
		t.Fatal(err)
	}

	if !s.Valid() {
		t.Fatalf("must not be null but got nil")
	}

	want := "foo"
	got := s.ToValue()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	err = json.NewDecoder(strings.NewReader("{}")).Decode(&s)
	if err == nil {
		t.Fatal("should be fail")
	}
}

func TestNullStringValueConverter(t *testing.T) {
	var s NullString

	err := s.Scan("1")
	if err != nil {
		t.Fatal(err)
	}

	if !s.Valid() {
		t.Fatalf("must not be null but got nil")
	}

	want := "1"
	got := s.ToValue()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	gotv, err := s.Value()
	if err != nil {
		t.Fatal(err)
	}
	if gotv != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

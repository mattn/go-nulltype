package nulltype

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestNullInt64Stringer(t *testing.T) {
	var b NullInt64

	want := "nil"
	got := fmt.Sprint(b)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = "3"
	b.Set(3)
	got = fmt.Sprint(b)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullInt64MarshalJSON(t *testing.T) {
	var b NullInt64

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(b)
	if err != nil {
		t.Fatal(err)
	}

	want := "null"
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	buf.Reset()

	b.Set(3)
	err = json.NewEncoder(&buf).Encode(b)
	if err != nil {
		t.Fatal(err)
	}

	want = "3"
	got = strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullInt64UnmarshalJSON(t *testing.T) {
	var b NullInt64

	err := json.NewDecoder(strings.NewReader("null")).Decode(&b)
	if err != nil {
		t.Fatal(err)
	}

	if b.Valid() {
		t.Fatalf("must be null but got %v", b)
	}

	err = json.NewDecoder(strings.NewReader(`3`)).Decode(&b)
	if err != nil {
		t.Fatal(err)
	}

	if !b.Valid() {
		t.Fatalf("must not be null but got nil")
	}

	want := int64(3)
	got := b.Value()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullInt64ValueConverter(t *testing.T) {
	var b NullInt64

	err := b.Scan("3")
	if err != nil {
		t.Fatal(err)
	}

	if !b.Valid() {
		t.Fatalf("must not be null but got nil")
	}

	want := int64(3)
	got := b.Value()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

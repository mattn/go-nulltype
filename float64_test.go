package nulltype

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestNullFloat64Stringer(t *testing.T) {
	var b NullFloat64

	want := "nil"
	got := fmt.Sprint(b)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = "3.14"
	b.Set(3.14)
	got = fmt.Sprint(b)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullFloat64MarshalJSON(t *testing.T) {
	var b NullFloat64

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

	b.Set(3.14)
	err = json.NewEncoder(&buf).Encode(b)
	if err != nil {
		t.Fatal(err)
	}

	want = "3.14"
	got = strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullFloat64UnmarshalJSON(t *testing.T) {
	var b NullFloat64

	err := json.NewDecoder(strings.NewReader("null")).Decode(&b)
	if err != nil {
		t.Fatal(err)
	}

	if b.Valid() {
		t.Fatalf("must be null but got %v", b)
	}

	b.Set(3.14)

	err = json.NewDecoder(strings.NewReader(`3.14`)).Decode(&b)
	if err != nil {
		t.Fatal(err)
	}

	if !b.Valid() {
		t.Fatalf("must not be null but got nil")
	}

	want := 3.14
	got := b.Value()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullFloat64ValueConverter(t *testing.T) {
	var b NullFloat64

	err := b.Scan("3.14")
	if err != nil {
		t.Fatal(err)
	}

	if !b.Valid() {
		t.Fatalf("must not be null but got nil")
	}

	want := 3.14
	got := b.Value()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

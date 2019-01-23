package nulltype

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestNullFloat64Stringer(t *testing.T) {
	var f NullFloat64

	want := "nil"
	got := fmt.Sprint(f)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = "3.14"
	f.Set(3.14)
	got = fmt.Sprint(f)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = "nil"
	f.Reset()
	got = fmt.Sprint(f)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullFloat64MarshalJSON(t *testing.T) {
	var f NullFloat64

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(f)
	if err != nil {
		t.Fatal(err)
	}

	want := "null"
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	buf.Reset()

	f.Set(3.14)
	err = json.NewEncoder(&buf).Encode(f)
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
	var f NullFloat64

	err := json.NewDecoder(strings.NewReader("null")).Decode(&f)
	if err != nil {
		t.Fatal(err)
	}

	if f.Valid() {
		t.Fatalf("must be null but got %v", f)
	}

	f.Set(3.14)

	err = json.NewDecoder(strings.NewReader(`3.14`)).Decode(&f)
	if err != nil {
		t.Fatal(err)
	}

	if !f.Valid() {
		t.Fatalf("must not be null but got nil")
	}

	want := 3.14
	got := f.Value()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullFloat64ValueConverter(t *testing.T) {
	var f NullFloat64

	err := f.Scan("3.14")
	if err != nil {
		t.Fatal(err)
	}

	if !f.Valid() {
		t.Fatalf("must not be null but got nil")
	}

	want := 3.14
	got := f.Value()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

package nulltype

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestNullStringStringer(t *testing.T) {
	var b NullString

	want := "nil"
	got := fmt.Sprint(b)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = "foo"
	b.Set("foo")
	got = fmt.Sprint(b)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullStringMarshalJSON(t *testing.T) {
	var b NullString

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

	b.Set("foo")
	err = json.NewEncoder(&buf).Encode(b)
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
	var b NullString

	err := json.NewDecoder(strings.NewReader("null")).Decode(&b)
	if err != nil {
		t.Fatal(err)
	}

	if b.Valid() {
		t.Fatalf("must be null but got %v", b)
	}

	err = json.NewDecoder(strings.NewReader(`"foo"`)).Decode(&b)
	if err != nil {
		t.Fatal(err)
	}

	if !b.Valid() {
		t.Fatalf("must not be null but got nil")
	}

	want := "foo"
	got := b.Value()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullStringValueConverter(t *testing.T) {
	var b NullString

	err := b.Scan("1")
	if err != nil {
		t.Fatal(err)
	}

	if !b.Valid() {
		t.Fatalf("must not be null but got nil")
	}

	want := "1"
	got := b.Value()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

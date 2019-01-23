package nulltype

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestNullBoolStringer(t *testing.T) {
	var b NullBool

	want := "nil"
	got := fmt.Sprint(b)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = "true"
	b.Set(true)
	got = fmt.Sprint(b)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = "false"
	b.Set(false)
	got = fmt.Sprint(b)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullBoolMarshalJSON(t *testing.T) {
	var b NullBool

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

	b.Set(true)
	err = json.NewEncoder(&buf).Encode(b)
	if err != nil {
		t.Fatal(err)
	}

	want = "true"
	got = strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	buf.Reset()

	b.Set(false)
	err = json.NewEncoder(&buf).Encode(b)
	if err != nil {
		t.Fatal(err)
	}

	want = "false"
	got = strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullBoolUnmarshalJSON(t *testing.T) {
	var b NullBool

	err := json.NewDecoder(strings.NewReader("null")).Decode(&b)
	if err != nil {
		t.Fatal(err)
	}

	if b.Valid() {
		t.Fatalf("must be null but got %v", b)
	}

	err = json.NewDecoder(strings.NewReader("true")).Decode(&b)
	if err != nil {
		t.Fatal(err)
	}

	if !b.Valid() {
		t.Fatalf("must not be null but got nil")
	}

	want := true
	got := b.Value()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullBoolValueConverter(t *testing.T) {
	var b NullBool

	err := b.Scan("1")
	if err != nil {
		t.Fatal(err)
	}

	if !b.Valid() {
		t.Fatalf("must not be null but got nil")
	}

	want := true
	got := b.Value()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

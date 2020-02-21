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

	want := ""
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
	b = NullBoolOf(false)
	got = fmt.Sprint(b)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = ""
	b.Reset()
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
	got := b.BoolValue()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	err = json.NewDecoder(strings.NewReader("false")).Decode(&b)
	if err != nil {
		t.Fatal(err)
	}

	if !b.Valid() {
		t.Fatalf("must not be null but got nil")
	}

	want = false
	got = b.BoolValue()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	err = json.NewDecoder(strings.NewReader(`"foo"`)).Decode(&b)
	if err == nil {
		t.Fatal("should be fail")
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
	got := b.BoolValue()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	gotv, err := b.Value()
	if err != nil {
		t.Fatal(err)
	}
	if gotv != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	b.Reset()

	gotv, err = b.Value()
	if err != nil {
		t.Fatal(err)
	}
	if gotv != nil {
		t.Fatalf("must be null but got %v", gotv)
	}
}

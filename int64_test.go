package nulltype

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestNullInt64Stringer(t *testing.T) {
	var i NullInt64

	want := "nil"
	got := fmt.Sprint(i)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = "3"
	i.Set(3)
	got = fmt.Sprint(i)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = "5"
	i = NullInt64Of(5)
	got = fmt.Sprint(i)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
	want = "nil"
	i.Reset()
	got = fmt.Sprint(i)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullInt64MarshalJSON(t *testing.T) {
	var i NullInt64

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(i)
	if err != nil {
		t.Fatal(err)
	}

	want := "null"
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	buf.Reset()

	i.Set(3)
	err = json.NewEncoder(&buf).Encode(i)
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
	var i NullInt64

	err := json.NewDecoder(strings.NewReader("null")).Decode(&i)
	if err != nil {
		t.Fatal(err)
	}

	if i.Valid() {
		t.Fatalf("must be null but got %v", i)
	}

	err = json.NewDecoder(strings.NewReader(`3`)).Decode(&i)
	if err != nil {
		t.Fatal(err)
	}

	if !i.Valid() {
		t.Fatalf("must not be null but got nil")
	}

	want := int64(3)
	got := i.ToValue()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	err = json.NewDecoder(strings.NewReader(`"foo"`)).Decode(&i)
	if err == nil {
		t.Fatal("should be fail")
	}
}

func TestNullInt64ValueConverter(t *testing.T) {
	var i NullInt64

	err := i.Scan("3")
	if err != nil {
		t.Fatal(err)
	}

	if !i.Valid() {
		t.Fatalf("must not be null but got nil")
	}

	want := int64(3)
	got := i.ToValue()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	gotv, err := i.Value()
	if err != nil {
		t.Fatal(err)
	}
	if gotv != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

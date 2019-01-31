package nulltype

import (
	"bytes"
	"html/template"
	"testing"
)

func TestTemplate(t *testing.T) {
	value := struct {
		Data1 string
		Data2 NullString
	}{}
	tpl, err := template.New("mytemplate").Parse(`{{.Data1}},{{.Data2}}`)
	if err != nil {
		t.Fatal(err)
	}
	var buf bytes.Buffer

	err = tpl.Execute(&buf, value)
	if err != nil {
		t.Fatal(err)
	}
	want := ","
	got := buf.String()
	if got != want {
		t.Fatalf("want %q, but %q:", want, got)
	}

	buf.Reset()
	value.Data1 = "data1"
	err = tpl.Execute(&buf, value)
	if err != nil {
		t.Fatal(err)
	}
	want = "data1,"
	got = buf.String()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	buf.Reset()
	value.Data2.Set("data2")
	err = tpl.Execute(&buf, value)
	if err != nil {
		t.Fatal(err)
	}
	want = "data1,data2"
	got = buf.String()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

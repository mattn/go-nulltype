package nulltype_test

import (
	"fmt"

	"github.com/mattn/go-nulltype"
)

type User struct {
	Name nulltype.NullString `json:"name"`
}

func Example_FmtStringer() {
	var user User
	fmt.Printf("%v, %q\n", user.Name.Valid(), user.Name)

	user.Name.Set("Bob")
	fmt.Printf("%v, %q\n", user.Name.Valid(), user.Name)

	fmt.Println(user.Name.StringValue() == "Bob") // true

	user.Name.Reset()
	fmt.Printf("%v, %q\n", user.Name.Valid(), user.Name)

	// Output:
	// false, ""
	// true, "Bob"
	// true
	// false, ""
}

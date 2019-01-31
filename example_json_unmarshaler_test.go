package nulltype_test

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Example_jsonUnmarshaler() {
	var user User
	s := `{"name": "Bob"}`
	json.NewDecoder(strings.NewReader(s)).Decode(&user)
	fmt.Printf("%v, %q\n", user.Name.Valid(), user.Name)

	s = `{"name": null}`
	json.NewDecoder(strings.NewReader(s)).Decode(&user)
	fmt.Printf("%v, %q\n", user.Name.Valid(), user.Name)

	// Output:
	// true, "Bob"
	// false, ""
}

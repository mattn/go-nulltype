package nulltype_test

import (
	"encoding/json"
	"fmt"
	"os"
)

func Example_jsonMarshaler() {
	var user User
	fmt.Println(user.Name.Valid())
	json.NewEncoder(os.Stdout).Encode(user)

	user.Name.Set("Bob")
	fmt.Println(user.Name.Valid())
	json.NewEncoder(os.Stdout).Encode(user)

	// Output:
	// false
	// {"name":null}
	// true
	// {"name":"Bob"}
}

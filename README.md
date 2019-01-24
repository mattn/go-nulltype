# go-nulltype

[![Build Status](https://travis-ci.org/mattn/go-nulltype.svg?branch=master)](https://travis-ci.org/mattn/go-nulltype)
[![codecov](https://codecov.io/gh/mattn/go-nulltype/branch/master/graph/badge.svg)](https://codecov.io/gh/mattn/go-nulltype)

Nullable types friendly to json.Encoder, json.Decoder, database/sql, fmt.Stringer, some of ORMs.

Supported types:

* NullBool
* NullString
* NullFloat64
* NullInt64
* NullTime

## Usage

```go
import "github.com/mattn/go-nulltype"

type User struct {
	Name	nulltype.NullString `json:"name"`
}
```

### friendly to Stringer

```go
var user User
fmt.Println(user.Name) // nil

user.Name.Set("Bob")
fmt.Println(user.Name) // Bob

fmt.Println(user.Name.Value() == "Bob") // true

user.Name.Reset()
fmt.Println(user.Name) // nil
```

### friendly to json.MarshalJSON

```go
var user User
json.NewEncoder(os.Stdout).Encode(user) // {"name": nil}

user.Name.Set("Bob")
json.NewEncoder(os.Stdout).Encode(user) // {"name": "Bob"}
```

### friendly to json.UnmarshalJSON

```go
var user User
s := `{"name": "Bob"}`
json.NewDecoder(strings.NewReader(s)).Decode(&user)
fmt.Println(user.Name) // Bob

s = `{"name": null}`
json.NewDecoder(strings.NewReader(s)).Decode(&user)
fmt.Println(user.Name) // nil
```

### friendly to database/sql

```go
var user User
db.QueryRow(`SELECT name FROM users`).Scan(&user.Name)
fmt.Println(user.Name) // Bob or nil
db.Exec(`INSERT INTO users(name) VALUES($1)`, user.Name)
```

### friendly to ORM

Struct tag with [gorp](https://github.com/go-gorp/gorp).

```go
type Post struct {
	Id      int64 `db:"post_id"`
	Created int64
	Title   string              `db:",size:50"`
	Body    nulltype.NullString `db:"body,size:1024"`
}
```

```go
p := Post{
	Created: time.Now().UnixNano(),
	Title:   title,
	Body:    nulltype.NullStringOf(body),
}
err = dbmap.Insert(&p)
```

## Installation

```
go get github.com/mattn/go-nullable
```

## License

MIT

## Author

Yasuhiro Matsumoto

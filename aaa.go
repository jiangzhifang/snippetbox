package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

	dsn := "dbname=snippetbox sslmode=disable user=web password=pass host=127.0.0.1"

	db, err := sql.Open("postgres", dsn)
	fmt.Println(err)

	email := "bob@xxx.com"

	a := db.QueryRow("SELECT id, hashed_password FROM users WHERE email = $1 AND active = TRUE", email)

	fmt.Println(a)

	var id int
	var hashedPassword []byte

	err = a.Scan(&id, &hashedPassword)
	// fmt.Println(err)
	fmt.Println(id)
	fmt.Println(string(hashedPassword))

	fmt.Println(err)

	em := err.Error()
	fmt.Printf("%s\n", em)

}

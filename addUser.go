package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	username string
	email    string
	password string
}

func addUser() {
	db, err := sql.Open("mysql", "root:bPqwkdR9hM56J4R8ggkV@tcp(217.76.55.135:3306)/magic")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var u User
	fmt.Print("Bitte Username eingeben:")
	fmt.Scan(&u.username)
	fmt.Print("Bitte Email eingeben:")
	fmt.Scan(&u.email)
	fmt.Print("password eingeben")
	fmt.Scan(&u.password)

	_, err = db.Exec("INSERT INTO User(username, email, password) VALUES(?, ?, ?)", u.username, u.email, u.password)
	if err != nil {
		panic(err.Error())
	}

	createTable := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s_collection (
            set_name VARCHAR(50) NOT NULL,
			set_num int NOT NULL,
            language VARCHAR(50) NOT NULL,
            foil BOOLEAN NOT NULL,
            count int Not NULL,
			deck_id int,
			box_id int,
			PRIMARY KEY (set_name, set_num, language)
        );
    `, u.username)
	_, err = db.Exec(createTable)
	if err != nil {
		panic(err.Error())
	}
}

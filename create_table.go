package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func createTable(tableName string) {
	db, err := sql.Open("mysql", "root:bPqwkdR9hM56J4R8ggkV@tcp(217.76.55.135:3306)/magic")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	createTable := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(50) NOT NULL,
            email VARCHAR(50) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    `, tableName)
	_, err = db.Exec(createTable)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Table created successfully!")
}

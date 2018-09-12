package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type Connection struct {
	db *sql.DB
}

func NewConnection() Connection {
	uri := os.Getenv("MYSQL_URI")

	if uri == "" {
		log.Fatal("MySQL enviroment url required.\neg `MYSQL_URI='root:password@tcp(127.0.0.1:3306)/'`")
	}

	db, err := sql.Open("mysql", uri)

	if err != nil {
		log.Fatal(err)
	}

	return Connection{db: db}
}

func (c *Connection) List(query string) {
	var rowData string

	rows, err := c.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&rowData)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(rowData)
	}
}

func (c Connection) Execute(cmd string) {
	_, err := c.db.Exec(cmd)

	if err != nil {
		log.Fatal(err)
	}
}

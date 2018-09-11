package main

import "fmt"

type Database struct {
	conn Connection
}

func (d Database) List() {
	d.conn.List("SHOW DATABASES;")
}

func (d Database) Add(database string) {
	d.conn.Execute(fmt.Sprintf("CREATE DATABASE `%s`;", database))
}

func (d Database) Drop(database string) {
	d.conn.Execute(fmt.Sprintf("DROP DATABASE `%s`;", database))
}

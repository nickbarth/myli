package main

import "fmt"

type Table struct {
	conn Connection
}

func (u Table) List(database string) {
	u.conn.Execute(fmt.Sprintf("USE `%s`;", database))
	u.conn.List("SHOW TABLES;")
}

func (u Table) Drop(database string, table string) {
	u.conn.Execute(fmt.Sprintf("USE `%s`;", database))
	u.conn.Execute(fmt.Sprintf("DROP TABLE `%s`;", table))
}

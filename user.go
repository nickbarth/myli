package main

import "fmt"

type User struct {
	conn Connection
}

func (u User) List() {
	u.conn.List("SELECT User FROM mysql.user;")
}

func (u User) Add(username string, password string, database string) {
	u.conn.Execute(fmt.Sprintf("CREATE USER `%s`@`%%` IDENTIFIED BY '%s';", username, password))
	u.conn.Execute(fmt.Sprintf("GRANT ALL PRIVILEGES ON `%s`.* TO `%s`@`%%`;", database, username))
	u.conn.Execute("FLUSH PRIVILEGES;")
}

func (u User) Drop(user string) {
	u.conn.Execute(fmt.Sprintf("DROP USER `%s`@`%%`;", user))
}

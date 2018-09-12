package main

import (
	"fmt"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"syscall"
)

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	app := cli.NewApp()

	app.Name = "myli"
	app.Usage = "mysql admin utils"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:  "user",
			Usage: "ls, add, rm",
			Subcommands: []cli.Command{
				{
					Name:  "ls",
					Usage: "list all users",
					Action: func(c *cli.Context) error {
						conn := NewConnection()
						user := User{conn: conn}
						user.List()
						return nil
					},
				},
				{
					Name:  "add",
					Usage: "[username] [database] -- add a new user",
					Action: func(c *cli.Context) error {
						username := c.Args().Get(0)
						database := c.Args().Get(1)

						if username == "" || database == "" {
							log.Fatal("myli: args required\nusage: myli user add [username] [database]")
						}

						fmt.Print("Enter Password: ")
						bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))

						if err != nil {
							log.Fatal(err)
						}

						conn := NewConnection()
						user := User{conn: conn}
						user.Add(username, string(bytePassword), database)

						fmt.Printf("\nuser `%s` created for `%s`.\n", username, database)

						return nil
					},
				},
				{
					Name:  "rm",
					Usage: "[username] -- remove a user",
					Action: func(c *cli.Context) error {
						username := c.Args().Get(0)

						if username == "" {
							log.Fatal("myli: args required\nusage: myli user rm [username]")
						}

						conn := NewConnection()
						user := User{conn: conn}
						user.Drop(username)

						fmt.Printf("user `%s` removed.\n", username)
						return nil
					},
				},
			},
		},
		{
			Name:  "db",
			Usage: "ls, add, rm",
			Subcommands: []cli.Command{
				{
					Name:  "ls",
					Usage: "list all databases",
					Action: func(c *cli.Context) error {
						conn := NewConnection()
						db := Database{conn: conn}
						db.List()
						return nil
					},
				},
				{
					Name:  "add",
					Usage: "[database] -- add a new database",
					Action: func(c *cli.Context) error {
						database := c.Args().Get(0)

						if database == "" {
							log.Fatal("myli: args required\nusage: myli db add [database]")
						}

						conn := NewConnection()
						db := Database{conn: conn}
						db.Add(database)

						fmt.Printf("database `%s` created.\n", database)

						return nil
					},
				},
				{
					Name:  "rm",
					Usage: "[database] -- remove a database",
					Action: func(c *cli.Context) error {
						database := c.Args().Get(0)

						if database == "" {
							log.Fatal("myli: args required\nusage: myli db rm [database]")
						}

						conn := NewConnection()
						db := Database{conn: conn}
						db.Drop(database)

						fmt.Printf("database `%s` removed.\n", database)
						return nil
					},
				},
			},
		},
		{
			Name:  "table",
			Usage: "ls, rm",
			Subcommands: []cli.Command{
				{
					Name:  "ls",
					Usage: "[database] -- list all tables",
					Action: func(c *cli.Context) error {
						database := c.Args().Get(0)

						if database == "" {
							log.Fatal("myli: args required\nusage: myli table ls [database]")
						}

						conn := NewConnection()
						db := Table{conn: conn}
						db.List(database)
						return nil
					},
				},
				{
					Name:  "rm",
					Usage: "[database] [table] -- remove a table",
					Action: func(c *cli.Context) error {
						database := c.Args().Get(0)
						table := c.Args().Get(1)

						if database == "" || table == "" {
							log.Fatal("myli: args required\nusage: myli table rm [database] [table]")
						}

						conn := NewConnection()
						db := Table{conn: conn}
						db.Drop(database, table)

						fmt.Printf("table `%s` removed.\n", table)
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

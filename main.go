package main

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	"github.com/gofiber/fiber"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Result struct {
	Id      int64
	Name    string
	Link    string
	Details string
	Types   string
	Tags    string
}

func main() {
	// fmt.Println("main.go")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	// Database connection
	db, err = sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}

	app := fiber.New()

	//Routes

	//GET Requests
	//Retrieves last row in the database
	app.Get("/api", func(c *fiber.Ctx) {
		query := "SELECT * FROM links ORDER BY id DESC LIMIT 1"

		res, err := db.Query(query)
		defer res.Close()
		if err != nil {
			log.Fatal("(GetLinks) db.Query", err)
		}

		results := []Result{}
		for res.Next() {
			var link Result
			err := res.Scan(&link.Id, &link.Name, &link.Link, &link.Details, &link.Types, &link.Tags)
			if err != nil {
				log.Fatal("(GetLinks) res.Scan", err)
			}
			results = append(results, link)
		}
		c.Send(results)

	})
	//GET Requests using parameters and SQL flexible search
	app.Get("/api/:id", func(c *fiber.Ctx) {
		//Search parameter
		searchValue := c.Params("id")
		query := fmt.Sprintf(`SELECT * FROM links WHERE tags LIKE %s%s%s`, "'%", searchValue, "%'")
		// fmt.Println(query)
		res, err := db.Query(query)
		defer res.Close()
		if err != nil {
			log.Fatal("(There was an error) db.Query", err)
		}

		//List of results
		results := []Result{}
		for res.Next() {
			var link Result
			err := res.Scan(&link.Id, &link.Name, &link.Link, &link.Details, &link.Types, &link.Tags)
			if err != nil {
				log.Fatal("(No results) res.Scan", err)
			}
			results = append(results, link)
		}
		c.Send(results)
	})

	app.Listen(3000)
}

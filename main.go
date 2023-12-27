package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"ps_go/schema"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)

var db *sql.DB

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

	//Routes
	app := fiber.New()
	//GET Requests
	app.Get("/api", func(c *fiber.Ctx) {
		//Retrieves last row in the database
		query := "SELECT * FROM links ORDER BY id DESC LIMIT 1"

		res, err := db.Query(query)
		defer res.Close()
		if err != nil {
			log.Fatal("(GetLinks) db.Query", err)
		}

		results := []schema.Result{}
		for res.Next() {
			var link schema.Result
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
		results := []schema.Result{}
		for res.Next() {
			var link schema.Result
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

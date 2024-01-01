package main

import (
	"database/sql"
	"log"

	"os"
	"ps_go/schema"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	// Open a connection to the database
	db, err = sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping: %v line 19", err)
	}
	log.Println("Successfully connected to PlanetScale!")

	app := fiber.New()

	//Routes
	app.Get("/api", func(c *fiber.Ctx) error {
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
		return c.JSON(results)
	})

	//Get route using parameters
	app.Get("/api/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		query := "SELECT * FROM links WHERE id = ? "

		res, err := db.Query(query, id)
		defer res.Close()
		if err != nil {
			log.Fatal("Problem with query", err)
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
		return c.JSON(results)
	})

	//Create and insert into database
	app.Post("/api", func(c *fiber.Ctx) error {
		link := new(schema.Result)

		if err := c.BodyParser(link); err != nil {
			log.Fatal("Unable to insert into db ", err)
		}

		query := `INSERT INTO links 
		(id, name, link, details, type, tags) VALUES (?, ?, ?, ?, ?, ?)`
		res, err := db.Exec(query, link.Id, link.Name, link.Link, link.Details, link.Types, link.Tags)
		if err != nil {
			log.Fatal("SQL error ", err)
		}

		link.Id, err = res.LastInsertId()
		if err != nil {
			log.Fatal("Error, res.LastInsertId ", err)
		}

		return c.JSON(link)
	})

	app.Listen(":3000")
}

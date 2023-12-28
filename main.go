package main

import (
	"database/sql"
	"ps_go/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"

	// "fmt"
	"log"
	// "os"
	// "ps_go/schema"
)

var db *sql.DB

func main() {
	//Database connection
	config.Config()

	//Routes
	app := fiber.New()
	//GET Requests
	app.Get("/api", func(c *fiber.Ctx) error {
		query := "SELECT * FROM links WHERE id > 160"
		res, err := db.Query(query)
		defer res.Close()
		if err != nil {
			log.Fatal("There was a problem db.Query", err)
		}
		defer res.Close()
	})

	app.Listen(":3000")
}

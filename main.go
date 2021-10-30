package main

import (
	"log"

	"github.com/ananyap/namedeelukrak/handlers"
	"github.com/ananyap/namedeelukrak/repositories"
	"github.com/ananyap/namedeelukrak/services"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:IntelliP24@tcp(localhost:3306)/ananyadb?parseTime=true")
	if err != nil {
		panic(err)
	}

	numberRepo := repositories.NewRepositoryDB(db)
	numberService := services.NewNumberService(numberRepo)
	numberHandler := handlers.NewNumberHandler(numberService)

	app := fiber.New()

	app.Get("/numbers", numberHandler.GetNumbers)

	log.Fatal(app.Listen(":3000"))

}

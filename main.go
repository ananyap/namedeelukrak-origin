package main

import (
	"fmt"
	"log"

	"github.com/ananyap/namedeelukrak/repositories"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := sqlx.Open("mysql", "root:IntelliP24@tcp(localhost:3306)/namedeelukrak?parseTime=true")
	if err != nil {
		panic(err)
	}

	//numberRepo := repositories.NewRepositoryDB(db)
	//numberService := services.NewNumberService(numberRepo)
	//numberHandler := handlers.NewNumberHandler(numberService)

	//number, err := numberService.GetNumber(3)

	userRepo := repositories.NewUserRepository(db)
	user, err := userRepo.Signup(repositories.UserSignupRequest{Username: "ananya", Passwd: "1234", Email: "ananya@email"})

	if err != nil {
		logrus.Error(err)
	}

	fmt.Println(user)

	app := fiber.New()

	//app.Get("/numbers", numberHandler.GetNumbers)
	//app.Get("/number/:id", numberHandler.GetNumber)

	log.Fatal(app.Listen(":3000"))

}

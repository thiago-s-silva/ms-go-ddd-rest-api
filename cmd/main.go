package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/thiago-s-silva/ms-go-task-management-api/infra/database"
	"github.com/thiago-s-silva/ms-go-task-management-api/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// init db
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	if db != nil {
		print("db ok")
	}

	// init api server
	e := echo.New()
	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

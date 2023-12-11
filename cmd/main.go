package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/thiago-s-silva/ms-go-task-management-api/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

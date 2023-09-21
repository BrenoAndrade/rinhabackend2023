package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/brenoandrade/rinhabackend2023/internal/people"
)

func main() {
	e := echo.New()

	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sqlx.Connect("postgres", dbURI)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	db.DB.SetMaxOpenConns(125)
	db.DB.SetConnMaxLifetime(0)

	people.RegisterModule(e.Group(""), db)

	e.Logger.Fatal(e.Start(":" + os.Getenv("SERVER_PORT")))
}

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sql_c/pkg/repository"
	"sql_c/pkg/seed"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
)

type Repo struct {
	db *repository.Queries
}

func dbConn() (*Repo, error) {
	viper.AutomaticEnv()
	// godotenv.Load(".env")

	ctx := context.Background()
	connStr := viper.GetString("DATABASE_URL")
	log.Println(connStr)
	if strings.TrimSpace(connStr) == "" {
		return nil, errors.New("DATABASE_URL is empty")
	}

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("db connection error: %w", err)
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("db ping error: %w", err)
	}

	return &Repo{
		db: repository.New(conn),
	}, nil
}

func main() {
	repo, err := dbConn()
	if err != nil {
		log.Fatalf("error: %v", err.Error())
	}
	seed.Seeder(repo.db)
	//repo.seed()
	app := fiber.New()

	app.Get("/users", repo.GetUsers)
	app.Post("/users", repo.CreateUser)
	app.Get("/user", repo.FindUser)

	if err := app.Listen(":8000"); err != nil {
		log.Fatalf("server error: %v", err.Error())
	}
}

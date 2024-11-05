package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDB() *pgxpool.Pool {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	dbpool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connected to the database.")
	return dbpool
}

// func ConnectDB() {
// 	conn, err := pgx.Connect(context.Background(), os.Gotenv("DATABASE_URL"))
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 		os.Exit(1)
// 	}
// 	defer conn.Close(context.Background())
// }

func CreateTableQuery(p *pgxpool.Pool) {
	_, err := p.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, full_name VARCHAR(255)), age INT, birth_date DATE, address VARCHAR(255))")
	if err != nil {
		log.Fatal("Error while creating table: ", err)
	}
}

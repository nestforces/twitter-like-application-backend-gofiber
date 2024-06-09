// database/database.go
package database

import (
    "context"
    "log"
    "twitter-like-backend/config"

    "github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func Connect() {
    var err error
    DB, err = pgxpool.Connect(context.Background(), config.AppConfig.DBUrl)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }

    log.Println("Connected to database")
}

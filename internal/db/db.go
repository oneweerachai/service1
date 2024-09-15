package db

import (
    "fmt"
    "github.com/oneweerachai/service1/internal/logger"

    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq" // Replace with your database driver
    "go.uber.org/zap"
)

var DB *sqlx.DB

func ConnectDB() (*sqlx.DB, error) {
    // Database configuration - in production, use environment variables or config files
    dbUser := "youruser"
    dbPassword := "yourpassword"
    dbName := "yourdb"
    dbHost := "localhost"
    dbPort := "5432"

    dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
        dbUser, dbPassword, dbName, dbHost, dbPort)

    db, err := sqlx.Connect("postgres", dsn)
    if err != nil {
        logger.NewLogger().Error("Database connection failed", zap.Error(err))
        return nil, err
    }

    // Optionally, set maximum connections, etc.
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(5 * 60)

    return db, nil
}

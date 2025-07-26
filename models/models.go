// Package models handles application configuration, including loading environment variables.
package models

import (
	"app/config"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	dialect := config.AppConfig.DBDialect
	username := config.AppConfig.DBUsername
	password := config.AppConfig.DBPassword
	host := config.AppConfig.DBHost
	port := config.AppConfig.DBPort
	debug := config.AppConfig.DBDebug == "YES"

	var dsn string
	var dialector gorm.Dialector

	switch dialect {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port)
		dialector = mysql.Open(dsn)

	case "mssql":
		dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%s", username, password, host, port)
		dialector = sqlserver.Open(dsn)

	case "postgres":
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, username, password)
		dialector = postgres.Open(dsn)

	default:
		err := fmt.Errorf("unsupported DB_DIALECT: %s", dialect)
		log.Println("Database connection failed:", err)
		return nil, fmt.Errorf("unsupported DB_DIALECT: %s", dialect)
	}

	logLevel := logger.Silent
	if debug {
		logLevel = logger.Info
	}

	var db *gorm.DB
	var err error

	for i := 1; i <= 5; i++ {
		db, err = gorm.Open(dialector, &gorm.Config{
			Logger: logger.Default.LogMode(logLevel),
		})
		if err == nil {
			break
		}
		log.Printf("Attempt %d: Database connection failed - %v", i, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Println("All attempts to connect to the database failed. Check username, password, host, and port.")
		return nil, fmt.Errorf("connection failed after retries: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get raw db error: %w", err)
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	log.Println("Database connection established successfully")
	DB = db
	return db, nil
}

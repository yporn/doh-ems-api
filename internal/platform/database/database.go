package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/yporn/doh-ems-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) {
	dbConfig := cfg.Database

	dsn := fmt.Sprintf(
		 "host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		 dbConfig.Host,
		 dbConfig.User,
		 dbConfig.Password,
		 dbConfig.Name,
		 dbConfig.Port,
		 dbConfig.SSLMode,
		dbConfig.TimeZone,
	)

	 fmt.Println("📦 DSN =>", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ failed to connect to database: %v", err)
	}

	DB = db
    fmt.Println("✅ PostgreSQL connected successfully!")
}

func ConnectMongo(mongoURI string) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("❌ failed to connect to MongoDB: %v", err)
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("❌ failed to ping MongoDB: %v", err)
		return nil, err
	}

	ParsedURL, err := connstring.ParseAndValidate(mongoURI)
	if err != nil {
		log.Fatalf("❌ failed to parse MongoDB URI: %v", err)
		return nil, err
	}

	dbName := ParsedURL.Database

	log.Println("✅ Successfully connected to MongoDB.")
	return client.Database(dbName), nil
}

package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Auction struct {
	ID            uint
	Quantity      uint
	InitialItemId uint
	UnitPrice     uint
	TimeLeft      string
	RunID         uint
}

type Run struct {
	ID        uint
	CreatedAt time.Time
}

func ConnectToDb() *gorm.DB {
	dsn := "host=localhost user=user password=example dbname=wow-auction-db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Issue connecting to database: %s\n", err)
		return nil
	}
	return db
}

func MigrateTables(db *gorm.DB) {
	err := db.AutoMigrate(&Run{}, &Auction{})
	if err != nil {
		fmt.Println("Error with migration", err)
	}
	db.Commit()
}

package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Auction struct {
	ID             uint
	Quantity       uint
	InitialItemId  uint
	UnitPrice      uint
	TimeLeft       string
	AuctionSliceID uint
}

type AuctionSlice struct {
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
	err := db.AutoMigrate(&AuctionSlice{}, &Auction{})
	if err != nil {
		fmt.Println("Error with migration", err)
	}
}

func InsertAuctionData(db *gorm.DB, auctionData *AuctionDataResponse) {
	auction_slice := AuctionSlice{}
	db.Create(&auction_slice)

	auctions := []Auction{}
	for _, data := range auctionData.Auctions {
		auction := Auction{Quantity: uint(data.Quantity), InitialItemId: uint(data.Item.ID), AuctionSliceID: auction_slice.ID, UnitPrice: uint(data.UnitPrice), TimeLeft: data.TimeLeft}
		auctions = append(auctions, auction)
	}
	fmt.Println("Start creation", time.Now())

	res := db.CreateInBatches(auctions, 100)
	fmt.Println("Committed", res.RowsAffected, res.Error, time.Now())
}

func retrieveLastSliceTime(db *gorm.DB) time.Time {
	var auction_slice AuctionSlice
	db.Order("created_at desc").Find(&auction_slice)
	return auction_slice.CreatedAt
}

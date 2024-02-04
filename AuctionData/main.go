package main

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

func getAuctionStuff() *AuctionDataResponse {
	auth := GetToken(retrieveEnvironmentValue("API_CLIENT_ID"), retrieveEnvironmentValue("API_CLIENT_SECRET"))
	auctionData := GetAuctionData(auth.AccessToken)
	return auctionData
}

func loadEnvironment() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
}

func retrieveEnvironmentValue(key string) string {
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

func main() {
	loadEnvironment()
	db := ConnectToDb()

	if db == nil {
		return
	}
	// MigrateTables(db)

	auctionData := getAuctionStuff()
	fmt.Println("Data Gathered", time.Now())
	run := Run{}
	runRes := db.Create(&run)
	fmt.Println("run created", runRes.RowsAffected, runRes.Error)

	auctions := []Auction{}
	for _, data := range auctionData.Auctions {
		auction := Auction{Quantity: uint(data.Quantity), InitialItemId: uint(data.Item.ID), RunID: run.ID, UnitPrice: uint(data.UnitPrice), TimeLeft: data.TimeLeft}
		auctions = append(auctions, auction)
	}
	fmt.Println("Start creation", time.Now())

	res := db.CreateInBatches(auctions, 100)
	fmt.Println("Committed", res.RowsAffected, res.Error, time.Now())

}

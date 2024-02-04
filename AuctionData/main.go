package main

import (
	"fmt"
	"time"
)

func getAuctionStuff() *AuctionDataResponse {
	auth := GetToken("38b0218c1fb24e4599b60c358b01964d", "tMzdyGzfOE01bCzc6yccCd186MK26TWG")
	auctionData := GetAuctionData(auth.AccessToken)
	return auctionData
}

func main() {
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

package main

import (
	"fmt"
)

func getAuctionStuff() {
	auth := GetToken("38b0218c1fb24e4599b60c358b01964d", "tMzdyGzfOE01bCzc6yccCd186MK26TWG")
	auctionData := GetAuctionData(auth.AccessToken)
	fmt.Println(auctionData.Auctions[0:10])
}

func main() {
	db := ConnectToDb()
	fmt.Println(db)

}

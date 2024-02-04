package main

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func getAuctionStuff(last_request_time time.Time) (auction_data *AuctionDataResponse, new_data bool) {
	auth := GetToken(retrieveEnvironmentValue("API_CLIENT_ID"), retrieveEnvironmentValue("API_CLIENT_SECRET"))
	auctionData, new_data := GetAuctionData(auth.AccessToken, last_request_time)
	return auctionData, new_data
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

func loadData(db *gorm.DB, last_request_time time.Time) {
	auctionData, new_data := getAuctionStuff(last_request_time)
	if !new_data {
		fmt.Println("No new data", time.Now())
		return
	}
	fmt.Println("Data Gathered", time.Now())
	InsertAuctionData(db, auctionData)
}

func main() {
	loadEnvironment()
	db := ConnectToDb()

	if db == nil {
		return
	}

	last_request_time := retrieveLastSliceTime(db)
	loadData(db, last_request_time)
}

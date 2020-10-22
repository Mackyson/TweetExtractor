package main

import (
	wrapper "TweetExtractor/internal/APIwrapper"
	textExtractor "TweetExtractor/internal/TextExtractor"
	"TweetExtractor/pkg/ESpkg"
	"TweetExtractor/pkg/Textpkg"
	"fmt"
	"log"
	"strings"
)

func main() {

	es, err := ESpkg.GetDBClient()
	if err != nil {
		log.Fatal(err)
	}

	tweets, err := wrapper.GetAllTweets(es, "restaurant")
	if err != nil {
		log.Fatal(err)
	}

	var restaurantNameList []string

	for _, tweet := range tweets {
		tweetText := tweet.Status["text"].(string)
		restaurantNameList = append(restaurantNameList, textExtractor.ExtractSpotName(tweetText))
	}
	resturantNameList := Textpkg.UniqueStrList(restaurantNameList)
	restaurantNameListCSV := strings.Join(resturantNameList, ",")
	fmt.Print(restaurantNameListCSV)
}

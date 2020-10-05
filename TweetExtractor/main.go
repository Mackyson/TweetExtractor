package main

import (
	wrapper "TweetExtractor/APIwrapper"
	"fmt"
	"log"
)

func main() {
	es, err := wrapper.GetDBClient()
	if err != nil {
		log.Fatal(err)
	}

	tweets, err := wrapper.GetAllTweets(es)
	if err != nil {
		log.Fatal(err)
	}
	for _, tweet := range tweets {
		fmt.Println("%+v", tweet.Urls)
	}
}

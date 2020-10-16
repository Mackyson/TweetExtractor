package main

import (
	wrapper "TweetExtractor/internal/APIwrapper"
	"TweetExtractor/pkg/ESpkg"
	"TweetExtractor/pkg/Textpkg"
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {

	op := flag.String("index", "", "index name (default: restaurant)")
	flag.Parse()

	es, err := ESpkg.GetDBClient()
	if err != nil {
		log.Fatal(err)
	}

	index := "restaurant"
	if *op != "" {
		index = *op
	}
	tweets, err := wrapper.GetAllTweets(es, index)
	if err != nil {
		log.Fatal(err)
	}
	var idList []string
	for _, tweet := range tweets {
		tweetUserId := tweet.Status["user"].(map[string]interface{})["id_str"].(string)
		idList = append(idList, tweetUserId)
	}
	idList = Textpkg.UniqueStrList(idList)
	idListCSV := strings.Join(idList, ",")
	fmt.Print(idListCSV)
}

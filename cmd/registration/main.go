package main

import (
	wrapper "TweetExtractor/internal/APIwrapper"
	textExtractor "TweetExtractor/internal/TextExtractor"
	"TweetExtractor/pkg/ESpkg"
	"TweetExtractor/pkg/Textpkg"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	stdin := bufio.NewScanner(os.Stdin)

	es, err := ESpkg.GetDBClient()
	if err != nil {
		log.Fatal(err)
	}

	tweets, err := wrapper.GetAllTweets(es, "tweet")
	if err != nil {
		log.Fatal(err)
	}

	for _, tweet := range tweets {
		tweetText := tweet.Status["text"].(string)
		tweetId := tweet.Status["id_str"].(string)
		isRegistered, err := wrapper.CheckRegistered(es, tweetId)
		if err != nil {
			log.Fatal(err)
		}
		if !isRegistered {
			//チェックインツイートのリツイートを拾っていた場合
			isRT := Textpkg.CheckRT(tweetText)
			if isRT {
				continue
			}
			spotName := textExtractor.ExtractSpotName(tweetText)
			log.Println(spotName)
			//restaurantインデックスに登録済みのSpotだった場合
			hasSameSpotDataInRestaurant, err := wrapper.CheckSameSpotDataWithIndex(es, "restaurant", spotName)
			if err != nil {
				log.Fatal(err)
			}
			if hasSameSpotDataInRestaurant {
				wrapper.RegisterAsRestaurant(es, tweet)
				continue
			}
			//otherインデックスに登録済みのSpotだった場合
			hasSameSpotDataInOther, err := wrapper.CheckSameSpotDataWithIndex(es, "other", spotName)
			if err != nil {
				log.Fatal(err)
			}
			if hasSameSpotDataInOther {
				wrapper.RegisterAsOther(es, tweet)
				continue
			}
			//その他の場合
			fmt.Printf("%s", tweetText)
			stdin.Scan()
			isNewRestaurant := stdin.Text() == "y"
			if isNewRestaurant {
				err = wrapper.RegisterAsRestaurant(es, tweet)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				err = wrapper.RegisterAsOther(es, tweet)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

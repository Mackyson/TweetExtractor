package main

import (
	wrapper "TweetExtractor/APIwrapper"
	textExtractor "TweetExtractor/TextExtractor"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	stdin := bufio.NewScanner(os.Stdin)

	es, err := wrapper.GetDBClient()
	if err != nil {
		log.Fatal(err)
	}

	tweets, err := wrapper.GetAllTweets(es)
	if err != nil {
		log.Fatal(err)
	}
	for _, tweet := range tweets {
		// isRegistered := false //データが入っていない状態で登録済みチェックインを探すとエラーを吐くので，そのときはコメントアウトを外す
		isRegistered, err := wrapper.CheckRegistered(es, tweet.Id)
		if err != nil {
			log.Fatal(err)
		}
		if !isRegistered {
			//チェックインツイートのリツイートを拾っていた場合
			isRT := textExtractor.CheckWhetherRT(tweet.Text)
			if isRT {
				continue
			}
			spotName := textExtractor.ExtractSpotName(tweet.Text)
			log.Println(spotName)
			//restaurantインデックスに登録済みのSpotだった場合
			hasSameSpotDataInRestaurant, err := wrapper.CheckSameSpotDataInRestaurant(es, spotName)
			if err != nil {
				log.Fatal(err)
			}
			if hasSameSpotDataInRestaurant {
				wrapper.RegisterAsRestaurant(es, tweet)
				continue
			}
			//otherインデックスに登録済みのSpotだった場合
			hasSameSpotDataInOther, err := wrapper.CheckSameSpotDataInOther(es, spotName)
			if err != nil {
				log.Fatal(err)
			}
			if hasSameSpotDataInOther {
				wrapper.RegisterAsOther(es, tweet)
				continue
			}
			//その他の場合
			fmt.Printf("%s", tweet.Text)
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

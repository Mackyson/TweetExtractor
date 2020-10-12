package main

import (
	wrapper "TweetExtractor/APIwrapper"
	textExtractor "TweetExtractor/TextExtractor"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	mode := flag.String("mode", "", "search query")
	op := flag.String("op", "", "option")
	flag.Parse()

	if *mode == "" {
		fmt.Println("Specify the mode and op")
		return
	}

	stdin := bufio.NewScanner(os.Stdin)

	es, err := wrapper.GetDBClient()
	if err != nil {
		log.Fatal(err)
	}

	tweets, err := wrapper.GetAllTweets(es)
	if err != nil {
		log.Fatal(err)
	}

	if *mode == "register" {

		for _, tweet := range tweets {
			tweetText := tweet.Status["text"].(string)
			tweetId := tweet.Status["id_str"].(string)
			isRegistered := false
			if *op != "new" {
				isRegistered, err = wrapper.CheckRegistered(es, tweetId)
			}
			if err != nil {
				log.Fatal(err)
			}
			if !isRegistered {
				//チェックインツイートのリツイートを拾っていた場合
				isRT := textExtractor.CheckWhetherRT(tweetText)
				if isRT {
					continue
				}
				spotName := textExtractor.ExtractSpotName(tweetText)
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
}

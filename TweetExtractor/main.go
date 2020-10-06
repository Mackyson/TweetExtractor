package main

import (
	wrapper "TweetExtractor/APIwrapper"
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
		// isRegistered := false //データが入っていない状態で登録済みチェックインを探すとエラーを吐くので，そのとき用
		isRegistered, err := wrapper.CheckRegistered(es, tweet.Id)
		if err != nil {
			log.Fatal(err)
		}
		if !isRegistered {
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

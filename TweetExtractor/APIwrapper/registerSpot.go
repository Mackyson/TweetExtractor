package APIwrapper

import (
	"TweetExtractor/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"strings"
)

func RegisterAsRestaurant(es *elasticsearch.Client, tweet model.Tweet) error {
	tweetInfoJSONText, err := json.Marshal(tweet)
	if err != nil {
		return err
	}
	req := esapi.IndexRequest{
		Index: "restaurant",
		Body:  strings.NewReader(string(tweetInfoJSONText)),
	}
	res, err := req.Do(context.Background(), es)
	res.Body.Close()
	if err != nil {
		return err
	}
	fmt.Println("stored as a restaurant")
	return nil
}

func RegisterAsOther(es *elasticsearch.Client, tweet model.Tweet) error {
	tweetInfoJSONText, err := json.Marshal(tweet)
	if err != nil {
		return err
	}
	req := esapi.IndexRequest{
		Index: "other",
		Body:  strings.NewReader(string(tweetInfoJSONText)),
	}
	res, err := req.Do(context.Background(), es)
	res.Body.Close()
	if err != nil {
		return err
	}
	fmt.Println("stored as NOT a restaurant")
	return nil
}
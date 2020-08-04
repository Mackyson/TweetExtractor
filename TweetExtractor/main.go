package main

import (
	// "context"
	// "encoding/json"
	"TweetExtractor/handler"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	// "github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	// "strings"
	// "time"
)

func GetDBClient() (*elasticsearch.Client, error) {
	es, err := elasticsearch.NewDefaultClient()
	return es, err
}

func main() {
	es, err := GetDBClient()
	if err != nil {
		log.Fatal(err)
	}

	tweets, err := handler.GetAllTweets(es)
	if err != nil {
		log.Fatal(err)
	}
	for _, tweet := range tweets {
		fmt.Println(tweet.UserId)
	}
	// query := "{\"query\": {\"match_all\": {}}, \"_source\": [\"user.id_str\",\"text\"]}"
	// req := esapi.SearchRequest{
	// 	Index:  []string{"tweet"},
	// 	Body:   strings.NewReader(query),
	// 	Scroll: 10 * time.Second,
	// }
	// res, err := req.Do(context.Background(), es)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var resJSON map[string]interface{}
	// for {
	// 	if err := json.NewDecoder(res.Body).Decode(&resJSON); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	res.Body.Close()
	// 	scrollID := resJSON["_scroll_id"].(string)
	// 	scrollReq := esapi.ScrollRequest{
	// 		ScrollID: scrollID,
	// 		Scroll:   10 * time.Second,
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if len(resJSON["hits"].(map[string]interface{})["hits"].([]interface{})) == 0 {
	// 		break
	// 	}
	// 	for _, tweet := range resJSON["hits"].(map[string]interface{})["hits"].([]interface{}) {
	// 		fmt.Printf("id: %s,text: %s\n", tweet.(map[string]interface{})["_source"].(map[string]interface{})["user"].(map[string]interface{})["id_str"], tweet.(map[string]interface{})["_source"].(map[string]interface{})["text"])
	// 	}
	// 	res, err = scrollReq.Do(context.Background(), es)
	// 	time.Sleep(1 * time.Second)
	// }

}

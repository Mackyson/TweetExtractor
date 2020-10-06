package APIwrapper

import (
	"TweetExtractor/model"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"strings"
	"time"
)

func GetAllTweets(es *elasticsearch.Client) ([]model.Tweet, error) {

	var (
		tweets []model.Tweet
		err    error
	)

	query := "{\"query\": {\"match_all\": {}}, \"_source\": [\"id_str\",\"user.id_str\",\"text\",\"entities\"]}"
	req := esapi.SearchRequest{
		Index:  []string{"tweet"},
		Body:   strings.NewReader(query),
		Scroll: 10 * time.Second,
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		return tweets, err
	}
	var resJSON map[string]interface{}
	for {
		if err = json.NewDecoder(res.Body).Decode(&resJSON); err != nil {
			return tweets, err
		}
		res.Body.Close()
		scrollID := resJSON["_scroll_id"].(string)
		scrollReq := esapi.ScrollRequest{
			ScrollID: scrollID,
			Scroll:   10 * time.Second,
		}
		if err != nil {
			return tweets, err
		}
		if len(resJSON["hits"].(map[string]interface{})["hits"].([]interface{})) == 0 {
			break
		}
		for _, tweetJSON := range resJSON["hits"].(map[string]interface{})["hits"].([]interface{}) {
			tweet := model.Tweet{
				Id:     tweetJSON.(map[string]interface{})["_source"].(map[string]interface{})["id_str"].(string),
				UserId: tweetJSON.(map[string]interface{})["_source"].(map[string]interface{})["user"].(map[string]interface{})["id_str"].(string),
				Text:   tweetJSON.(map[string]interface{})["_source"].(map[string]interface{})["text"].(string),
			}
			for _, urlJSON := range tweetJSON.(map[string]interface{})["_source"].(map[string]interface{})["entities"].(map[string]interface{})["urls"].([]interface{}) {
				tweet.Urls = append(tweet.Urls, urlJSON.(map[string]interface{})["expanded_url"].(string))
			}
			//map[string]interface{}へのキャストを繰り返して，JSONをちょっとずつパースしている．
			tweets = append(tweets, tweet)
		}
		res, err = scrollReq.Do(context.Background(), es)
		if err != nil {
			return tweets, err
		}
		time.Sleep(time.Millisecond * 1)
	}
	return tweets, err
}

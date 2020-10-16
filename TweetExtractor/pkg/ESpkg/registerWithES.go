package ESpkg

import (
	"TweetExtractor/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"strings"
)

func RegisterWithIndex(es *elasticsearch.Client, indexName string, searchResponse model.SearchResponse) error {
	tweetInfoJSONText, err := json.Marshal(searchResponse.Status)
	if err != nil {
		return err
	}
	req := esapi.IndexRequest{
		Index: indexName,
		Body:  strings.NewReader(string(tweetInfoJSONText)),
	}
	res, err := req.Do(context.Background(), es)
	res.Body.Close()
	if err != nil {
		return err
	}
	fmt.Println("stored in the " + indexName + " index")
	return nil
}

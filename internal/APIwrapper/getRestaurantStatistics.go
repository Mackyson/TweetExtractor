package APIwrapper

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io"
	// "os"
	"strings"
	"time"
)

func GetRestaurantStatistics(es *elasticsearch.Client, indexName string, query string) (map[string]int, error) { //あとでまともな名前にする

	var (
		err error
	)
	count := make(map[string]int)
	req := esapi.SearchRequest{
		Index:  []string{indexName},
		Body:   strings.NewReader(query),
		Scroll: 100 * time.Millisecond,
	}
	res, err := req.Do(context.Background(), es)
	var r io.Reader = res.Body
	// r = io.TeeReader(r, os.Stderr)
	if err != nil {
		return count, err
	}
	var resJSON map[string]interface{}
	for {
		if err = json.NewDecoder(r).Decode(&resJSON); err != nil {
			return count, err
		}
		res.Body.Close()
		scrollID := resJSON["_scroll_id"].(string)
		scrollReq := esapi.ScrollRequest{
			ScrollID: scrollID,
			Scroll:   100 * time.Millisecond,
		}
		clearScrollReq := esapi.ClearScrollRequest{
			ScrollID: []string{scrollID},
		}
		if err != nil {
			return count, err
		}
		if len(resJSON["hits"].(map[string]interface{})["hits"].([]interface{})) == 0 {
			clearScrollReq.Do(context.Background(), es)
			break
		}
		for _, tweetJSON := range resJSON["hits"].(map[string]interface{})["hits"].([]interface{}) {
			id := tweetJSON.(map[string]interface{})["_source"].(map[string]interface{})["user"].(map[string]interface{})["id_str"].(string)
			count[id] += 1
		}
		res, err = scrollReq.Do(context.Background(), es)
		r = res.Body
		// r = io.TeeReader(r, os.Stderr)

		if err != nil {
			return count, err
		}

		time.Sleep(time.Millisecond * 1)
	}
	time.Sleep(time.Millisecond * 1)
	return count, err
}

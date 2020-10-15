package APIwrapper

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	dbg "log"
	"strings"
)

func CheckSameSpotDataWithIndex(es *elasticsearch.Client, indexName string, spotName string) (bool, error) {
	query := "{\"query\":{\"bool\":{\"should\":[{\"match_phrase\":{\"text\":\"at " + spotName + " in 高松市 \"}},{\"match_phrase\":{\"text\":\"at " + spotName + " in Takamatsu\"}}]}}}"
	req := esapi.SearchRequest{
		Index: []string{indexName},
		Body:  strings.NewReader(query),
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		return false, err
	}
	var resJSON map[string]interface{}
	if err = json.NewDecoder(res.Body).Decode(&resJSON); err != nil {
		dbg.Printf("%+v", resJSON)
		return false, err
	}
	res.Body.Close()
	if len(resJSON["hits"].(map[string]interface{})["hits"].([]interface{})) == 0 {
		return false, nil
	} else {
		return true, nil
	}
	return false, err
}

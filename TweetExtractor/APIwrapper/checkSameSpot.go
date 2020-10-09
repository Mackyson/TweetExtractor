package APIwrapper

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	dbg "log"
	"strings"
)

//TODO クエリの結果を取得する部分は切り分けてもいいかも
func CheckSameSpotDataInRestaurant(es *elasticsearch.Client, spotName string) (bool, error) {
	query := "{\"query\": {\"match_phrase\": {\"text\":\" at " + spotName + " in 高松市\"}}}"
	req := esapi.SearchRequest{
		Index: []string{"restaurant"},
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

	// dbg.Printf("%+v", resJSON)
	if len(resJSON["hits"].(map[string]interface{})["hits"].([]interface{})) == 0 {
		return false, nil
	} else {
		return true, nil
	}
	return false, err
}
func CheckSameSpotDataInOther(es *elasticsearch.Client, spotName string) (bool, error) {
	query := "{\"query\": {\"match_phrase\": {\"text\":\"" + spotName + " in 高松市\"}}}"
	req := esapi.SearchRequest{
		Index: []string{"other"},
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

	// dbg.Printf("%+v", resJSON)
	if len(resJSON["hits"].(map[string]interface{})["hits"].([]interface{})) == 0 {
		return false, nil
	} else {
		return true, nil
	}
	return false, err
}

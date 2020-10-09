package APIwrapper

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	dbg "log"
	"strings"
)

func CheckRegistered(es *elasticsearch.Client, id string) (bool, error) {
	query := "{\"query\": {\"term\": {\"id_str\":" + id + "}}}"
	req := esapi.SearchRequest{
		Index: []string{"restaurant", "other"},
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

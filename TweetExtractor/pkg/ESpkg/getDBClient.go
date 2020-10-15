package ESpkg

import "github.com/elastic/go-elasticsearch/v8"

func GetDBClient() (*elasticsearch.Client, error) {
	es, err := elasticsearch.NewDefaultClient()
	return es, err
}

package APIwrapper

import (
	"TweetExtractor/model"
	"TweetExtractor/pkg/ESpkg"

	"github.com/elastic/go-elasticsearch/v8"
)

func RegisterAsRestaurant(es *elasticsearch.Client, searchResponse model.SearchResponse) error {
	err := ESpkg.RegisterWithIndex(es, "restaurant", searchResponse)
	return err
}

func RegisterAsOther(es *elasticsearch.Client, searchResponse model.SearchResponse) error {
	err := ESpkg.RegisterWithIndex(es, "other", searchResponse)
	return err
}

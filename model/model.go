package model

type SearchResponse struct {
	Status map[string]interface{} `json:"status"`
}

type Tweet struct /*FIXME:別の名前を募集中*/ {
	Id     string   `json:"id_str"`
	UserId string   `json:"user.id_str"`
	Text   string   `json:"text"`
	Urls   []string `json:"urls"`
}

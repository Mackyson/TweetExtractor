package model

type Tweet struct /*FIXME:別の名前を募集中*/ {
	UserId string `json:"user.id_str"`
	Text   string `json:"text"`
}

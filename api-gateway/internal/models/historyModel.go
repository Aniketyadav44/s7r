package models

import "time"

type History struct {
	Id        string    `json:"id"`
	UrlId     string    `json:"url_id"`
	ClickedAt time.Time `json:"clicked_at"`
}

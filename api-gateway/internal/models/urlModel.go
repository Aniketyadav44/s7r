package models

import "time"

type URL struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Url       string    `json:"url"`
	Short     string    `json:"short"`
	Clicks    int       `json:"clicks"`
	Expiry    time.Time `json:"expiry"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

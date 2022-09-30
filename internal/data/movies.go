package data

import "time"

type Movie struct {
	PublicID  int64     `json:"id,omitempty"`
	PrivateID int64     `json:"-"`
	Title     string    `json:"title"`
	Genres    []string  `json:"genres"`
	Year      int32     `json:"year"`
	Runtime   int32     `json:"runtime_in_mins,omitempty,string"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

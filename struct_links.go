package shikimori

import "time"

type Links struct {
	Id         int       `json:"id"`
	Kind       string    `json:"kind"`
	Url        string    `json:"urls.go"`
	Source     string    `json:"source"`
	EntryId    int       `json:"entry_id"`
	EntryType  string    `json:"entry_type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ImportedAt time.Time `json:"imported_at"`
}

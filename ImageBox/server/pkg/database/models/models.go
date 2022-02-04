package models

import "time"

type Annotation struct {
	ID             int64     `json:"id"`
	CollectionName string    `json:"collectionName"`
	Filename       string    `json:"filename"`
	Description    string    `json:"description"`
	Filepath       string    `json:"filepath"`
	Favourite      bool      `json:"favourite"`
	Created        time.Time `json:"created"`
}

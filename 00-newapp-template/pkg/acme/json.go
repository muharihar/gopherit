package acme

import (
	"encoding/json"
)

type Gopher struct {
	ID          json.Number `json:"gopher_id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
}

type Thing struct {
	ID          json.Number `json:"thing_id"`
	GopherID    json.Number `json:"gopher_id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
}

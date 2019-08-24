package extranormal

import "time"

type Item struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Class       string    `json:"class"`
	Value       float64   `json:"value"`
	Weight      float64   `json:"weight"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type ItemService interface {
	CreateItem() (*Item, error)
}

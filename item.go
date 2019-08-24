package extranormal

type Item struct {
	Name        string  `json:"name"`
	Class       string  `json:"class"`
	Value       float64 `json:"value"`
	Weight      float64 `json:"weight"`
	Description string  `json:"description"`
}

type ItemService interface {
	CreateItem() (*Item, error)
}
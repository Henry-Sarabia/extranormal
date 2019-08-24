package craft

import "github.com/Henry-Sarabia/extranormal"

type ItemService struct {
	gen *Generator
}

func NewItemService(g *Generator) *ItemService {
	return &ItemService{gen: g}
}

func (s *ItemService) CreateItem() (*extranormal.Item, error) {
	return s.gen.Item()
}
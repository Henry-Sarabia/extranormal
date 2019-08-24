package craft

import (
	"github.com/Henry-Sarabia/extranormal"
	"github.com/Henry-Sarabia/placeholder"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Generator struct {
	gen *placeholder.Crafter
}

func NewGenerator() *Generator {
	return &Generator{
		gen: &placeholder.Crafter{},
	}
}

func (g *Generator) Item() (*extranormal.Item, error) {
	i, err := g.gen.NewItem()
	if err != nil {
		return nil, errors.Wrap(err, "generator cannot get new item")
	}

	id, err := uuid.NewV4()
	if err != nil {
		return nil, errors.Wrap(err, "cannot generate uuid")
	}

	return &extranormal.Item{
		ID:          id.String(),
		Name:        i.Name,
		Class:       i.Class,
		Value:       i.Value,
		Weight:      i.Weight,
		Description: i.Description,
		CreatedAt:   time.Now(),
	}, nil
}

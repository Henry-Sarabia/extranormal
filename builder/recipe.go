package builder

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
)

type Recipe struct {
	Name       string       `json:"name"`
	BaseValue  float64      `json:"base_value"`
	BaseWeight float64      `json:"base_weight"`
	Comps      []Component `json:"components"`
}

func readRecipes(r ...io.Reader) ([]Recipe, error) {
	if len(r) <= 0 {
		return nil, errors.New("cannot read Recipes without at least one io.Reader")
	}

	var rec []Recipe

	for i, v := range r {
		temp, err := readRecipe(v)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot decode Recipes from io.Reader at index %d", i)
		}

		rec = append(rec, temp...)
	}

	return rec, nil
}

func readRecipe(r io.Reader) ([]Recipe, error) {
	var rec []Recipe

	if err := json.NewDecoder(r).Decode(&rec); err != nil {
		return nil, errors.Wrap(err, "cannot decode Recipe from io.Reader")
	}

	return rec, nil
}
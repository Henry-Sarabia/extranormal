package builder

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
)

type Attribute struct {
	Name               string   `json:"name"`
	WeightFactor       factor   `json:"weight_factor"`
	MinorValueFactor   factor   `json:"minor_value_factor"`
	MinorValueVariants []string `json:"minor_value_variants"`
	AvgValueFactor     factor   `json:"avg_value_factor"`
	AvgValueVariants   []string `json:"avg_value_variants"`
	MajorValueFactor   factor   `json:"major_value_factor"`
	MajorValueVariants []string `json:"major_value_variants"`
	Prefixes           []string `json:"prefix_references"`
}

func readAttributes(r ...io.Reader) ([]Attribute, error) {
	if len(r) <= 0 {
		return nil, errors.New("cannot read Attributes without at least one io.Reader")
	}

	var attr []Attribute

	for i, v := range r {
		temp, err := readAttribute(v)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot decode Attributes from io.Reader at index %d", i)
		}

		attr = append(attr, temp...)
	}

	return attr, nil
}

func readAttribute (r io.Reader) ([]Attribute, error) {
	var attr []Attribute

	if err := json.NewDecoder(r).Decode(&attr); err != nil {
		return nil, errors.Wrap(err, "cannot decode Attribute from io.Reader")
	}

	return attr, nil
}
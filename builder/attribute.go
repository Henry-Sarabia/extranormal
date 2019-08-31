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

// readAttribute reads the JSON-encoded Attributes from the provided Reader.
func readAttribute (r io.Reader) ([]Attribute, error) {
	var attr []Attribute

	if err := json.NewDecoder(r).Decode(&attr); err != nil {
		return nil, errors.Wrap(err, "cannot decode Attribute from io.Reader")
	}

	return attr, nil
}
package builder

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
)

type AttributeGroup struct {
	Name           string   `json:"name"`
	AttributeNames []string `json:"attribute_names"`
	Attributes     []*Attribute
}

// ReadAttributeGroups reads the JSON-encoded AttributeGroups from the provided
// Readers.
func ReadAttributeGroups(r ...io.Reader) ([]AttributeGroup, error) {
	if len(r) <= 0 {
		return nil, errors.New("cannot read AttributeGroups without at least one io.Reader")
	}

	var attr []AttributeGroup

	for i, v := range r {
		temp, err := readAttributeGroup(v)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot decode AttributeGroups from io.Reader at index %d", i)
		}

		attr = append(attr, temp...)
	}

	return attr, nil
}

// readAttributeGroup reads the JSON-encoded AttributeGroups from the provided
// Reader.
func readAttributeGroup (r io.Reader) ([]AttributeGroup, error) {
	var attr []AttributeGroup

	if err := json.NewDecoder(r).Decode(&attr); err != nil {
		return nil, errors.Wrap(err, "cannot decode AttributeGroup from io.Reader")
	}

	return attr, nil
}
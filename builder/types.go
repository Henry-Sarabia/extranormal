package builder

type Recipe struct {
	Name       string       `json:"name"`
	BaseValue  float64      `json:"base_value"`
	BaseWeight float64      `json:"base_weight"`
	Comps      []*Component `json:"components"`
}

type Component struct {
	Name       string      `json:"name"`
	Required   bool        `json:"required"`
	Properties []*Property `json:"properties"`
}

type Property struct {
	Name                string   `json:"name"`
	Required            bool     `json:"required"`
	AttributeNames      []string `json:"attribute_names"`
	AttributeGroupNames []string `json:"attribute_group_names"`
	Attributes          []*Attribute
	AttributeGroups     []*AttributeGroup
}

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

type AttributeGroup struct {
	Name           string   `json:"name"`
	AttributeNames []string `json:"attribute_names"`
	Attributes     []*Attribute
}

type factor float64

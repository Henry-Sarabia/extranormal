package builder

type Property struct {
	Name                string   `json:"name"`
	Required            bool     `json:"required"`
	AttributeNames      []string `json:"attribute_names"`
	AttributeGroupNames []string `json:"attribute_group_names"`
	Attributes          []*Attribute
	AttributeGroups     []*AttributeGroup
}
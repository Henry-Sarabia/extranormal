package builder

type Item struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Class       string  `json:"class"`
	Value       float64 `json:"value"`
	Weight      float64 `json:"weight"`
	Description string  `json:"description"`
}

type ItemBuilder interface {
	SetRecipes([]*Recipe)
	SetAttributes([]*Attribute)
	SetAtrributeGroups([]*AttributeGroup)
	LinkResources() error
	Item() (*Item, error)
}

type MundaneBuilder struct {
	Recipes         map[string]*Recipe
	Attributes      map[string]*Attribute
	AttributeGroups map[string]*AttributeGroup
}

func (mb *MundaneBuilder) SetRecipes([]*Recipe) {

}

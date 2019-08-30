package builder

type Component struct {
	Name       string      `json:"name"`
	Required   bool        `json:"required"`
	Properties []Property `json:"properties"`
}
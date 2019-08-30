package builder

import "github.com/pkg/errors"

type MundaneBuilder struct {
	Recipes    map[string]*Recipe
	Attributes map[string]*Attribute
	Groups     map[string]*AttributeGroup
}

func (mb *MundaneBuilder) LinkResources() error {
	if err := mb.linkGroups(); err != nil {
		return errors.Wrap(err, "cannot link groups") //TODO: Elaborate on error message
	}

	if err := mb.linkRecipes(); err != nil {
		return errors.Wrap(err, "cannot link recipes") //TODO: Elaborate
	}

	return nil
}

// linkGroups iterates through every AttributeGroup's AttributeNames and adds
//the corresponding Attribute addresses to the AttributeGroup's Attributes slice.
func (mb *MundaneBuilder) linkGroups() error {
	for _, g := range mb.Groups {
		for _, name := range g.AttributeNames {
			attr, ok := mb.Attributes[name]
			if !ok {
				return errors.Errorf("cannot find Attribute '%s' from AttributeGroup '%s' in builder's loaded Attributes", name, g.Name)
			}

			g.Attributes = append(g.Attributes, attr)
		}
	}
	return nil
}

// linkRecipes links every Recipe's AttributeNames and AttributeGroupNames to
// to their respective Attributes and AttributeGroups.
func (mb *MundaneBuilder) linkRecipes() error {
	for _, rec := range mb.Recipes {
		for _, comp := range rec.Comps {
			for _, prop := range comp.Properties {
				for _, name := range prop.AttributeNames {
					attr, ok := mb.Attributes[name]
					if !ok {
						return errors.Errorf("cannot find Attribute '%s' from Recipe '%s' in builder's loaded Attributes", name, rec.Name)
					}

					prop.Attributes = append(prop.Attributes, attr)
				}

				for _, name := range prop.AttributeGroupNames {
					g, ok := mb.Groups[name]
					if !ok {
						return errors.Errorf("cannot find AttributeGroup '%s' from Recipe '%s' in builder's loaded AttributeGroups", name, rec.Name)
					}

					prop.AttributeGroups = append(prop.AttributeGroups, g)
				}
			}
		}
	}
	return nil
}

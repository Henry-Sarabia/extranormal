package main

import (
	"fmt"
	"github.com/Henry-Sarabia/extranormal/craft"
)

func main() {
	g := craft.NewGenerator()
	s := craft.NewItemService(g)

	fmt.Println(s.CreateItem())
}
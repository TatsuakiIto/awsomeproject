package main

import "fmt"

func main()  {
	type SaveProperty struct{
		Name *string
	}

	p1Name := "test001"
	p2Name := "test002"

	p1 := SaveProperty{Name: &p1Name}
	p2 := SaveProperty{Name: &p2Name}

	var properties []*SaveProperty
	properties = append(properties, &p1)
	properties = append(properties, &p2)

	fmt.Println(properties)
	
	fmt.Println(properties[0].Name)
}
package main

import (
	"fmt"
)

type Vertex2 struct {
	Lat, Long float64
}

//var m map[string]Vertex2

var m = map[string]Vertex2{
	"Bell Labs": Vertex2{
		40.68433, -74.39967,
	},
	"Google": Vertex2{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m) // nil map, cannot use it

	//m = make(map[string]Vertex2) //make initialize and ready for use


	fmt.Println(m)


	//insert element
	m["MercadoLibre"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m)

	//get element
	el, ok := m["MercadoLibre"]
	fmt.Println(el, "is_present: ", ok)

	//delete element
	delete(m, "MercadoLibre")
	fmt.Println(m)

	//Iterate
	for key, element := range m {
		if key == "Google" {
			element.Lat = 0000
		}
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

	fmt.Println(m)
}
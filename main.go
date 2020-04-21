package main

import (
	"fmt"

	"github.com/pluralsight/datatypes/organization"
)

func main() {
	p := organization.NewPerson("Paulo", "Langer", organization.NewEuropeanUnionIdentifier(123456789, "Germany"))

	err := p.SetTwitterHandler("@paulolanger")
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}

	println(p.ID())
	println(p.Country())
}

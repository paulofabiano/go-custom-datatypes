package main

import (
	"fmt"

	"github.com/pluralsight/datatypes/organization"
)

func main() {
	p := organization.NewPerson("Paulo", "Langer")

	err := p.SetTwitterHandler("@paulolanger")
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}

	println(p.ID())
	println(p.FullName())
	println(p.TwitterHandler())
	println(p.TwitterHandler().RedirectUrl())
}

package main

import "github.com/pluralsight/datatypes/organization"

func main() {
	p := organization.NewPerson("James", "Wilson")
	println(p.ID())
	println(p.FullName())
}

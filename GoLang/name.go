package main

import "fmt"

type personsNames []string

func newPerson() personsNames {
	names := personsNames{"Ralph", "Mike"}
	return names
}

func (p personsNames) print() {
	for _, name := range p {
		fmt.Println(name)
	}
}

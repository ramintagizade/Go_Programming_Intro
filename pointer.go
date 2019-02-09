package main

import (
	"fmt"
)

func zero(x int ) {
	x = 0
}

func zeroPointer(x *int) {
	*x = 0
}

type Person struct {
	name string
	nextPerson *Person
}

func main(){

	x :=5	
	y:= 10
	zeroPointer(&x)
	zeroPointer(&y)
	fmt.Println(x)

	first_person := Person{}
	first_person.name = "Ramin"
	second_person := Person{}
	second_person.name = "Taghizada"
	first_person.nextPerson = &second_person
	fmt.Println(first_person.name)
	fmt.Println(second_person.name)
}
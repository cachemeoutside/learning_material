package main

import "fmt"

type dog struct {
	name  string
	breed string
	age   int
}

func newDog(name string, breed string, age int) *dog {
	d := dog{name, breed, age}
	return &d
}

func main() {
  d0 := dog{"random", "Aussie", 0}
	d1 := newDog("Mocha", "Aussie", 1)
	fmt.Printf("%s is an %s and is %v years old\n", d1.name, d1.breed, d1.age)
	fmt.Printf("%s is an %s and is %v years old\n", d0.name, d0.breed, d0.age)
}

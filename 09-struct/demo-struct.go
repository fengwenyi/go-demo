package main

import "fmt"

func main() {
	dog := Dog{
		Animal{
			"小狗",
		},
	}
	dog.eat()
}

type Animal struct {
	Name string
}

type Dog struct {
	Animal
}

func (a Animal) eat() {
	fmt.Printf("%v 吃东西\n", a.Name)
}

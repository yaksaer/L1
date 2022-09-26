package main

import (
	"fmt"
)

type Human struct {
	name string
	age  int
}

func (h Human) Say() {
	fmt.Printf("Hello, my name is %s I am %d years old\n", h.name, h.age)
}

type Action struct {
	Human
	action string
	name   string
}

func (a Action) doAction() {
	fmt.Printf("%s can %s\n", a.name, a.action)
	fmt.Printf("%s can %s\n", a.Human.name, a.Human)
}

func (a Action) Say() {
	fmt.Printf("Hello, my name is %s I am %d years old I can %s\n", a.name, a.age, a.action)
}

func main() {
	//Create struct Action
	act := Action{
		action: "shake hands",
		Human: Human{
			name: "Jhon",
			age:  35,
		},
		name: "Bob",
	}
	act2 := Action{
		action: "dance",
		Human: Human{
			name: "Alison",
			age:  20,
		},
		name: "Helen",
	}
	//Call Action's method
	act.doAction()
	act2.doAction()
	//Human's method call with shadowing
	act.Say()
	//Human's method call without shadowing
	act.Human.Say()
	//Human's method call with shadowing
	act2.Say()
	//Human's method call without shadowing
	act2.Human.Say()
}

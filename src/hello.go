package main

import "fmt"

type IAttack interface {
	Attack()
}

type Warrior struct {
	Name string
}

type Empire struct {
	Name string
}

func (w Warrior) Attack() {
	fmt.Printf("%s attacks with a sword!\n", w.Name)
}

func (e Empire) Attack() {
	fmt.Printf("%s attacks with a legion!\n", e.Name)
}

func main() {
	var attacker IAttack
	attacker = &Warrior{Name: "Conan"}
	attacker.Attack()

	attacker = &Empire{Name: "Roman Empire"}
	attacker.Attack()
}

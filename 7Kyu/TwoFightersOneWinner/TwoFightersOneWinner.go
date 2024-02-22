package main

// Create a function that returns the name of the winner in a fight between two fighters.
// Each fighter takes turns attacking the other and whoever kills the other first is victorious.
// Death is defined as having health <= 0.
// Each fighter will be a Fighter object/instance.
// See the Fighter class below in your chosenlanguage.
// Both health and damagePerAttack (damage_per_attack for python) will be integers larger than 0.
// You can mutate the Fighter objects.
// Your function also receives a third argument, a string, with the name of the fighter that
// attacks first.

import "fmt"

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func main() {
	// Prints for test
	fmt.Println(DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Lew"))
	fmt.Println(DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Harry"))
	fmt.Println(DeclareWinner(Fighter{"Harald", 20, 5}, Fighter{"Harry", 5, 4}, "Harry"))
	fmt.Println(DeclareWinner(Fighter{"Harald", 20, 5}, Fighter{"Harry", 5, 4}, "Harald"))
	fmt.Println(DeclareWinner(Fighter{"Jerry", 30, 3}, Fighter{"Harald", 20, 5}, "Jerry"))
	fmt.Println(DeclareWinner(Fighter{"Jerry", 30, 3}, Fighter{"Harald", 20, 5}, "Harald"))
}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	var fAttack, sAttack *Fighter

	if firstAttacker == fighter1.Name {
		fAttack = &fighter1
		sAttack = &fighter2
	} else {
		fAttack = &fighter2
		sAttack = &fighter1
	}

	for fAttack.Health >= 0 || sAttack.Health >= 0{
		sAttack.Health -= fAttack.DamagePerAttack
		if sAttack.Health <= 0 { return fAttack.Name }
		
		fAttack.Health -= sAttack.DamagePerAttack
		if fAttack.Health <= 0 { return sAttack.Name }
	}

	return "Untouchable return OMEGALUL"
}

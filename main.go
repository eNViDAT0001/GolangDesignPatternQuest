package main

import (
	observerDesign "DesignPatternQuest/behavioural/observer"
	creationalDesign "DesignPatternQuest/creational"
	"fmt"
)

func main() {
	//Creational
	fmt.Println("Creational Pattern")
	//singleton: Create static instance
	fmt.Println("Singleton Pattern: ")
	counter := creationalDesign.GetInstance()
	fmt.Println(counter.AddOne())
	fmt.Printf("%p\n", counter)
	//builder: Create instance with parameter helper
	fmt.Println("Builder Pattern: ")
	account := creationalDesign.NewAccountBuilder().WithID(9999).WithBalance(9990).WithUserName("eNViDAT").Build()
	fmt.Println(account.GetInfo())
	//Factory: Hide logic behind
	physicAttack := creationalDesign.ChooseAttack(creationalDesign.Physic)
	physicAttack.Hit(20)
	spellAttack := creationalDesign.ChooseAttack(creationalDesign.Spell)
	spellAttack.Hit(20)

	//Behavioral
	fmt.Println("Behavioral Pattern")
	//Observer: Pub-Sub pattern
	channel := observerDesign.CreateChannel("Gud boy channel")
	bob := observerDesign.NewUser("Bob")
	peter := observerDesign.NewUser("Peter")
	channel.Register(bob)
	channel.Register(peter)
	channel.NotifyAll()
	//Iterator

	//State

	//Strategy

	//Structural
	fmt.Println("Structural Pattern")

	//adapter

	//bridge

	//composite

	//proxy

}

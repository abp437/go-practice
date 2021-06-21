package main

import (
	"fmt"
	"go-practice/langintro"
)

func sayGreeting(name string) {
	fmt.Println("Good morning", name)
}

func main() {
	langintro.PrintHelloWorld()
	langintro.PrintHelloNinjas()

	marioCharactersList := langintro.MarioCharacters()
	fmt.Println(marioCharactersList, len(marioCharactersList))

	langintro.NormalLoop(5)
	langintro.RangeLoop()

	langintro.HigherOrderFunction("Wario", sayGreeting)
	fmt.Println(langintro.CircleArea(34.6))
	fn, ln := langintro.PersonInitials("Akshay Pawar")
	fmt.Printf("%v. %v. \n", fn, ln)
	fmt.Println(langintro.MyName)

	fmt.Println(langintro.ModifyName("Akshay"))

	myBill := langintro.NewBill("Akshay")
	myBill.AddItem("Tea", 10)
	myBill.AddItem("Biscuits", 15)
	myBill.UpdateTip(24.05)
	fmt.Print(myBill.Format())
}

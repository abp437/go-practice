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
}

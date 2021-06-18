package main

import (
	"fmt"
	"go-practice/langintro"
)

func main() {
	langintro.PrintHelloWorld()
	langintro.PrintHelloNinjas()

	marioCharactersList := langintro.MarioCharacters()
	fmt.Println(marioCharactersList, len(marioCharactersList))
}

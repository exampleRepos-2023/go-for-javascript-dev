package main

import (
	"fmt"
)

var mySentence = "Neco naber"

func main() {
	for index, letter := range mySentence {
		fmt.Println("INDEX: ", index, " LETTER: ", string(letter))
	}
}

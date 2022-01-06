package main

import (
	"fmt"

	"github.com/SoliDeoHonertGloria/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println("found", word, "defination:", hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}

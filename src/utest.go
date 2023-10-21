package main

import (
	"fmt"
	"strings"
)

func main(){
	var acc Account
	var phrases []string = strings.Split("TEST PHRASES", " ")
	acc.fromPhrases(phrases)
	fmt.Println(acc.Seed)
	phrases = AccountGenerateSeed("../res/json/words.json", 15)
	if phrases!=nil{
		fmt.Println(strings.Join(phrases, " "))
	}
}
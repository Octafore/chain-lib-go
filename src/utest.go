package main

import (
	"fmt"
	"strings"
)

func main(){
	var acc Account
	var phrases []string = strings.Split("TEST PHRASES", " ")
	acc.fromPhrases(phrases)
	fmt.Print(acc.Seed)
}
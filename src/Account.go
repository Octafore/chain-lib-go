package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

type Account struct{
	Number int			`json:"number"`
	Seed string 		`json:"seed"`
	Phrases []string 	`json:"phrases"` 	
	Key	string			`json:"key"`
	PKey string 		`json:"pkey"`
	Algorithm string 	`json:"algorithm"`
}

func (acc *Account) fromPhrases(phrases []string){
	pstr := strings.Join(phrases, " ")
	hash := sha256.New()
	hash.Write([]byte(pstr))
	acc.Seed = fmt.Sprintf("%x", hash.Sum(nil))
}

func AccountGenerateSeed(word_file string, count uint) []string{
	if count==0 {
		count = 12
	}
	return nil
}
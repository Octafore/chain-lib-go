package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"math/rand"
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

func AccountGenerateSeed(word_file string, count int) []string{
	dat, err := os.ReadFile(word_file)
	if err==nil{
		var words []string
		json.Unmarshal(dat, &words)
		if count==0 {
			count = 12
		}
		var result = make([]string, count)
		i := 0
		for i<count {
			index := rand.Intn(len(words))
			word := words[index]
			if len(word)>3 && inArray(word, result)<0{
				result[i] = word
				i += 1
			}
		}
		return result[:]
	}else{
		fmt.Println(err)
	}
	return nil
}
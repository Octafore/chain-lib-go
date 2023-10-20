package main

type Account struct{
	Number int			`json:"number"`
	Seed string 		`json:"seed"`
	Key	string			`json:"key"`
	PKey string 		`json:"pkey"`
	Algorithm string 	`json:"algorithm"`
}
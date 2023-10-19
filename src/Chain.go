package main

type Chain struct{
	ID string 				`json:"id"`
	FeeRate string			`json:"feeRate"`
	MainToken Token 		`json:"mainToken"`
}
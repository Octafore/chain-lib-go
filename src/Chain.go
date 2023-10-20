package main

type Chain struct{
	ID string 				`json:"id"`
	FeeRate string			`json:"feeRate"`
	MainToken Token 		`json:"mainToken"`
	Epoch uint 				`json:"epoch"`
	Version float32			`json:"version"`
	SupportedAlgos []string `json:"supportedAlgos"`
}
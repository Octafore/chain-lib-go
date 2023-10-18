package main

type TX struct{
	Hash string					`json:"hash"`
	Block string				`json:"block"`
	Time int64					`json:"time"`
	Sender string				`json:"sender"`
	Inputs []interface{}		`json:"inputs"`
	Outputs []interface{}		`json:"outputs"`
	Fee	string					`json:"fee"`
}
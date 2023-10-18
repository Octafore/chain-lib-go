package main

type Block struct{
	Chain string			`json:"chain"`
	Height int64			`json:"heigh"`
	Time int64				`json:"time"`
	Hash string				`json:"hash"`
	Previous string			`json:"previous"`
	Seed string				`json:"seed"`
	Data []interface{}		`json:"data"`
	Options	int 			`json:"options"`
	Verifier string			`json:"verifier"`
}
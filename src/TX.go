package main

type TX struct{
	//block data
	Hash string					`json:"hash"`
	Model string 				`json:"model"`
	//body
	Block string				`json:"block"`
	Time uint64					`json:"time"`
	Sender string				`json:"sender"`
	Inputs []interface{}		`json:"inputs"`
	Outputs []interface{}		`json:"outputs"`
	Fee	string					`json:"fee"`
}

func TXCreate() TX{
	var ret TX
	ret.Model = "TX"
	ret.Hash = "Invalid"
	return ret
}
package main

type Block struct{
	Chain string			`json:"chain"`
	Epoch uint				`json:"epoch"`
	Version float32			`json:"version"`
	Height int64			`json:"heigh"`
	Time uint64				`json:"time"`
	Hash string				`json:"hash"`
	Previous string			`json:"previous"`
	Seed string				`json:"seed"`
	Data []interface{}		`json:"data"`
	Options	int 			`json:"options"`
	Finalizer string		`json:"finalizer"`
	Algorithm string 		`json:"algorithm"`
}

func BlockCreate(chain Chain) Block{
	var ret Block
	ret.Hash = "Invalid"
	ret.Chain = chain.ID
	ret.Version = chain.Version
	return ret
}
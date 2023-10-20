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
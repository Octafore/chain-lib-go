package main

// block data and it's children will store inside the block

type BlockData struct{
	Hash string 				`json:"hash"`
	Model string				`json:"model"`
	Algo string 				`json:"algo"`
}

func BlockDataCreate() BlockData{
	var ret BlockData
	ret.Model = "BlockData"
	ret.Hash = "Invalid"
	ret.Algo = "Default"
	return ret
}
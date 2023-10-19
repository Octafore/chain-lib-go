package main

type BlockData struct{
	Hash string 				`json:"hash"`
	Model string				`json:"model"`
}

func BlockDataCreate() BlockData{
	var ret BlockData
	ret.Model = "BlockData"
	ret.Hash = "Invalid"
	return ret
}
package main

// Network data and it's children will store inside network storage (if available on node)
// not inside blocks
// owner account is able to change/update the data

type NetData struct{
	ID string 					`json:"id"`
	Model string				`json:"model"`
	Account string 				`json:"account"`
	Signature string			`json:"signature"`
}

func NetDataCreate() NetData{
	var ret NetData
	ret.Model = "BlockData"
	ret.ID = "Invalid"
	ret.Signature = "Invalid"
	return ret
}
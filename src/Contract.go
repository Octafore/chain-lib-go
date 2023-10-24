package main
import "time"
type Contract struct{
	BlockData
	//
	Account string				`json:"account"`
	Time uint64					`json:"time"`
	VmName string				`json:"vmName"`
	Code string					`json:"code"`
	Source []interface{}		`json:"source"`
	Version float32 			`json:"version"`
}

func ContractCreate() Contract{
	var ret Contract
	ret.Model = "BlockData"
	ret.Hash = "Invalid"
	ret.Algo = "Default"

	ret.Account = "Invalid"
	ret.Time = uint64(time.Now().Unix())
	ret.VmName = "lua"
	ret.Code = ""
	ret.Version = 0.00

	return ret
}
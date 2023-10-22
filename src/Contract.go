package main

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
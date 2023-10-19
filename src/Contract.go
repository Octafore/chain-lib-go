package main

type Contract struct{
	//block data
	Hash string					`json:"hash"`
	Model string				`json:"model"`
	//
	Account string				`json:"account"`
	Time uint64					`json:"time"`
	VmName string				`json:"vmName"`
	Code string					`json:"code"`
	Source []interface{}		`json:"source"`
}
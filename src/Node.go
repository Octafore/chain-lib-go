package main

type Node struct{
	Account string 				`json:"account"`
	Endpoint string 			`json:"endpoint"`
	Type int 					`json:"type"`
	Chains []string 			`json:"chains"`
	Draft []interface{}			`json:"draft"`
	Inbox []interface{}			`json:"inbox"`
}

const NODE_TYPE_FINALIZER 		= 1
const NODE_TYPE_WALLET			= 2
const NODE_TYPE_VERIFIER		= 4
const NODE_TYPE_MASTER			= 255
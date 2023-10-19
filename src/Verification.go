package main

type Verification struct{
	//block data
	Hash string 			`json:"hash"`
	Model string 			`json:"model"`
	//body
	Node string 			`json:"node"`
	Account string 			`json:"account"`
	Type int 				`json:"type"`
	Time uint64 				`json:"time"`
	Signature string 		`json:"signature"`
}

const VERIFICATION_TYPE_ID				= 1
const VERIFICATION_TYPE_ADDRESS 		= 2
const VERIFICATION_TYPE_BUSINESS_DOCS	= 4

func VerificationCreate() Verification{
	var ret Verification
	ret.Hash = "Invalid"
	ret.Model = "Verification"
	return ret
}
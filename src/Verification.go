package main

import "time"

type Verification struct{
	NetData
	//body
	Node string 			`json:"node"`
	Type int 				`json:"type"`
	Created uint64			`json:"created"`
	Modified uint64			`json:"modified"`
	Signature string 		`json:"signature"`
}

const VERIFICATION_TYPE_ID				= 1
const VERIFICATION_TYPE_ADDRESS 		= 2
const VERIFICATION_TYPE_BUSINESS_DOCS	= 4

func VerificationCreate() Verification{
	var ret Verification
	ret.ID = "Invalid"
	ret.Model = "Verification"
	ret.Account = "Invalid"
	ret.Signature = "Invalid"
	ret.Created = uint64(time.Now().Unix()) 
	ret.Modified = ret.Created
	return ret
}
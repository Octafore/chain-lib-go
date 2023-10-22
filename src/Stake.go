package main

type Stake struct{
	BlockData
	//
	OwnerAccount string				`json:"ownerAccount"`
	RedeemAccount string			`json:"redeemAccount"`
	Node string						`json:"node"`
	InterestRate float32			`json:"interestRate"`
	Amount string 					`json:"amount"`
	Token string 					`json:"token"`
	Time uint64						`json:"time"`
	Delivery uint64 				`json:"delivery"`
	Lockdown uint64 				`json:"lockdown"`
	Options uint 					`json:"options"`
	OwnerSignature string 			`json:"ownerSignature"`
	NodeSignature string 			`json:"nodeSignature"`
}

const STAKE_OPT_AUTO_DELIVER	= 1;

func StakeCreate() Stake{
	var ret Stake
	ret.Model = "Stake"
	ret.Hash = "Invalid"
	return ret
}
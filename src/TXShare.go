package main

type TXShare struct{
	Tx string				`json:"tx"`
	Account string			`json:"account"`
	Amount string			`json:"amount"`
	Token string			`json:"token"`
	Option int				`json:"options"`
}

const TX_SHARE_LOCKED	= 1
const TX_SHARE_FEE 		= 2
const TX_SHARE_PAYED	= 4

func TXShareCreate() TXShare{
	var ret TXShare
	return ret
}
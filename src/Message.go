package main

type Message struct{
	NetData
	//
	Target string			`json:"target"`
	Text string				`json:"text"`
	Resouces []Resource		`json:"resources"`
	Created uint64			`json:"created"`
	Modified uint64			`json:"modified"`
}

func MessageCreate() Message{
	var ret Message
	ret.ID = "Invalid"
	ret.Model = "Resource"
	ret.Account = "Invalid"
	ret.Signature = "Invalid"
	return ret
}
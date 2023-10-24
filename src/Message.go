package main

import "time"
type Message struct{
	NetData
	//
	Target string			`json:"target"`
	Text string				`json:"text"`
	Format string			`json:"format"`   // plain | html | md
	Resouces []Resource		`json:"resources"`
	Created uint64			`json:"created"`
	Modified uint64			`json:"modified"`
}

func MessageCreate() Message{
	var ret Message

	ret.Model = "Message"
	ret.ID = "Invalid"
	ret.Account = "Invalid"
	ret.Signature = "Invalid"

	ret.Format = "plain"
	ret.Created = uint64(time.Now().Unix())
	ret.Modified = ret.Created
	
	return ret
}
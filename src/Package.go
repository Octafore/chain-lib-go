package main

type Package struct{
	Payload	[]interface{}			`json:"payload"`
	SenderID string 				`json:"senderId"`
	SenderSignature string 			`json:"senderSignature"`
}
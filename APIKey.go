package main

type APIKey struct{
	Title string 			`json:"title"`
	User string				`json:"user"`
	Key string 				`json:"key"`		//public key
	Secret string 			`json:"secret"`		//private key
	Permission int 			`json:"permissions"`
	Expiry uint64 			`json:"Expiry"`
}
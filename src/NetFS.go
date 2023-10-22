package main

type NetFS struct{
	NetData
	//
	MimeType string				`json:"mimeType"`
	// local file maybe different foreach node
	LocalFile string 			`json:"localFile"`
	//binary/text data to transfrer trough the network
	Data []byte 				`json:"data"`
}
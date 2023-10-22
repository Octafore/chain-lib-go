package main

type Resource struct{
	NetData
	//
	MimeType string			`json:"mimeType"`
	Endpoint string			`json:"endpoint"`
	Description string		`json:"description"`
}

func ResourceCreate() Resource{
	var ret Resource
	ret.ID = "Invalid"
	ret.Model = "Resource"
	ret.Account = "Invalid"
	ret.Signature = "Invalid"
	return ret
}
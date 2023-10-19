package main

type Token struct{
	Name string 				`json:"name"`
	Abbrevation string 			`json:"abbrevation"`
	Account string 				`json:"account"`
	TotalAmount string 			`json:"totalAmount"`
	Communities []string		`json:"communities"`
	Resources []interface{}		`json:"resources"`
	Contract string 			`json:"contract"` 
}
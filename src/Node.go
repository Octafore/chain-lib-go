package main

type Node struct{
	ID string 					`json:"id"`
	Account string 				`json:"account"`
	Key string 					`json:"key"`
	Endpoint string 			`json:"endpoint"`
	Type int 					`json:"type"`
	Chains []string 			`json:"chains"`
	SupportedAlgos []string 	`json:"supportedAlgos"`
	DefaultAlgo string 			`json:"defaultAlgo"`
	// not exporting fields
	pKey string //private key
	nodeList []Node
	parentNode interface{}
	childrenNodes []interface{}
	draft []interface{}
	inbox []interface{}
	sendObjectFunc func(pkg Package, node Node)  //setup sendObjectFunc inside application
}

const NODE_TYPE_FINALIZER 		= 1
const NODE_TYPE_WALLET			= 2
const NODE_TYPE_VERIFIER		= 4
const NODE_TYPE_MASTER			= 255


//
func (a *Node) SetSendFunction(fn func(pkg Package, node Node)){
	a.sendObjectFunc = fn
}
//
func (a *Node) GetNodeById(id string, list []Node) *Node{
	i := 0
	for i<len(a.nodeList){
		if(a.nodeList[i].ID==id){
			return &a.nodeList[i]
		}
	}
	return nil
}
func (n *Node) Income(obj interface{}) bool{
	return false
}
# Structure Models

- Account
- Node
- Block
- TX
- TX Share

---
## Package:
All the information from client app to high level node must transfer inside Package Object and json Format. Package Object consists of:

| Property | Description |
|----------|-------------|
| payload | array of Objects like TX, Block, Message .... |
| senderId | node Id or API Key  |
| senderSignature | signature of (payload json)  |


## Account:
Account class is use to provide access and keep information about an account and consists of following properties:

| Property | Description |
|----------|-------------|
| number | Account number |
| seed | seed used to generate private key |
| key | public key |
| pkey | private key |

## Node:
Node class is main container of every network node. it specifies the node behaviour and keep data in order:

| Property | Description |
|----------|-------------|
| account | Node Account to receive fees |
| endpoint | access the node through the network |
| type | one or more node type |
| chains | supported chain ids |
| draft | incomming data from other nodes |
| inbox | finailized data from other nodes  to store and indexing |

## Block Data:
Block data based objects are object those store inside block

| Property | Description |
|----------|-------------|
| hash | unique hash for the object |
| model | type name of the object like TX |
| algo | hashing (and signature) algorithm |


## NetData:
Net data base objects are objects those transfer through the network and store inside different storage type like NetFS or Database. they may be editable/removable by the owner or permission users but not store inside blocks itself. some of them may expire after specific time and remove from the whole network

## Block:
Block Structure used to store data in order and make the network decenteralized.

| Property | Description |
|----------|-------------|
| chain | keep chain to specify the block belongs to which chain |
| Epoch | block belongs to which blockchain epoch |
| Version | network version |
| height | block height is the block index starting from 0 |
| time | block close time |
| hash | block unique hash id |
| previous | previous block hash |
| seed | block seed can be a random number to guarantee block hash is unique |
| data | raw data to store in the block |
| options | used to specify the state of the block like: open, closed, .... |
| finalizer | the id of node who finalized and closed the block |


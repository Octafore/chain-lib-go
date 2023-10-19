# Structure Models

- Account
- Node
- Block
- TX
- TX Share

---

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

## Block:
Block Structure used to store data in order and make the network decenteralized.

| Property | Description |
|----------|-------------|
| chain | keep chain to specify the block belongs to which chain |
| height | block height is the block index starting from 0 |
| time | block close time |
| hash | block unique hash id |
| previous | previous block hash |
| seed | block seed can be a random number to guarantee block hash is unique |
| data | raw data to store in the block |
| options | used to specify the state of the block like: open, closed, .... |
| finalizer | the id of node who finalized and closed the block |


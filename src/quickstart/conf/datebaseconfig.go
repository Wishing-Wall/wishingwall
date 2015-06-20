/* package -- config  setting the configuration for source code
 *databaseConfig.go : setting the databases configuration
 *author: Johnson Bernoulli
 */
package conf

const DATABASEPATH string = "./database.db"

const BLOCK_FIRST uint64 = 1

//main net
//const WISHINGWALLADDRESS string = "DLHV2GJrDL5M9atZ49BZ6DKwZhDWFEZfxw"

//testnet
const WISHINGWALLADDRESS string = "mwn9HEbcjFuxZd9AZZZsLjPKVRhPxGEXV2"

//const PAYTOWISHINGWALL string = "OP_DUP OP_HASH160 b2616c2a516f75fa26896919cb52602338a71308 OP_EQUALVERIFY OP_CHECKSIG"
const PAYTOWISHINGWALL string = "76a914b2616c2a516f75fa26896919cb52602338a7130888ac"
const COIN uint64 = 100000000
const MESSAGEFEE = 10860

const FEE = 10000

const OP_1 = "51"
const OP_2 = "52"
const OP_CHECKMULTISIG = "ae"

//runtime variable
var CURRENT_BLOCK_INDEX uint64

//tables
type DB_message struct {
	Id            int
	Message_count uint64 //the message_count in tran.Data

	Block_index_list string
	Tx_index_list    string
	Tx_hash_list     string
	Account          uint64
	Source           string // a unique symbol for a DB_message
	Destination      string
	Message          string //the combine of message_body in many tran.Data
}
type DB_messages []DB_message

func (l DB_messages) Len() int {
	return len(l)
}
func (l DB_messages) Less(i, j int) bool {
	return l[i].Id < l[j].Id
}
func (l DB_messages) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type DB_transaction struct {
	Id          int
	Tx_index    uint64
	Tx_hash     string
	Block_index uint64
	Block_hash  string
	Block_time  uint64
	Source      string
	Destination string
	Btc_amount  uint64
	Fee         uint64
	Data        string
	Supported   bool
}

type DB_blocks struct {
	Id              int
	Block_index     uint64
	Block_hash      string
	Block_prev_hash string
	Block_time      uint64
}

type DB_send struct {
	Id           int
	RelayAddr    string
	Message      string
	ConfirmTimes int
	CheckTimes   int
	IsSent       bool
	Succeed      bool
}

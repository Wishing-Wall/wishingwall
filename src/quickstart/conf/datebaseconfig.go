/* package -- config  setting the configuration for source code
 *databaseConfig.go : setting the databases configuration
 *author: Johnson Bernoulli
 */
package conf

const DATABASEPATH string = "./database.db"

const BLOCK_FIRST uint64 = 1
const WISHINGWALLADDRESS string = "DLHV2GJrDL5M9atZ49BZ6DKwZhDWFEZfxw"

//runtime variable
var CURRENT_BLOCK_INDEX uint64

//tables
type DB_message struct {
	Id            int
	Message_index uint64
	Block_index   uint64
	Block_hash    string
	Block_time    uint64
	Tx_index      uint64
	Tx_hash       string
	Account       string
	Source        string
	Destination   string
	Message       string
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

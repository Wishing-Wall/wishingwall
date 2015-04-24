/* package -- config  setting the configuration for source code
 *databaseConfig.go : setting the databases configuration
 *author: Johnson Bernoulli
 */
package conf

const DATABASEPATH string = "./database.db"

const BLOCKFIRST uint64 = 1
const WISHINGWALLADDRESS string = "DLHV2GJrDL5M9atZ49BZ6DKwZhDWFEZfxw"

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

type DB_transaction struct {
	Id          int
	Tx_index    uint64
	Tx_hash     string
	Block_index int64
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

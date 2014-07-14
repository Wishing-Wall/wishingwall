/* package -- config  setting the configuration for source code
 *databaseConfig.go : setting the databases configuration
 *author: Johnson Bernoulli
 */
package config

const DATABASEPATH string = "./database/database.db"

const DATABASE_CREATE_BLOCK string = "CREATE TABLE IF NOT EXISTS blocks (" +
	"block_index INTEGER UNIQUE," +
	"block_hash TEXT UNIQUE," +
	"block_time INTEGER," +
	"PRIMARY KEY (block_index,block_hash))"
const DATABASE_CREATE_BLOCK_INDEX string = "CREATE INDEX IF NOT EXISTS " +
	"block_index_idx ON blocks (block_index)"
const DATABASE_CREATE_BLOCK_HASH_INDEX string = "CREATE INDEX IF NOT EXISTS " +
	"index_hash_idx ON blocks (block_index, block_hash)"

const DATABASE_TRANSACTION string = "CREATE TABLE IF NOT EXISTS transactions (" +
	"tx_index INTEGER UNIQUE," +
	"tx_hash TEXT UNIQUE," +
	"block_index INTEGER," +
	"block_hash TEXT," +
	"block_time INTERGER," +
	"source TEXT," +
	"destination TEXT," +
	"btc_amount INTEGER," +
	"fee INTERGER," +
	"data BLOB," +
	"supported BOOL DEFAULT 1," +
	"FOREIGN KEY (block_index, block_hash) REFERENCES blocks(block_index,block_hash)," +
	"PRIMARY KEY (tx_index, tx_hash,block_index))"
const DATABASE_CREATE_TRANSACTION_INDEX string = "CREATE INDEX IF NOT EXISTS " +
	"block_index_idx ON transactions (block_index)"
const DATABASE_CREATE_TRANSACTION_INDEX2 string = "CREATE INDEX IF NOT EXISTS " +
	"tx_index_idx ON transactions (tx_index)"
const DATABASE_CREATE_TRANSACTION_INDEX3 string = "CREATE INDEX IF NOT EXISTS " +
	"tx_hash_idx ON transactions (tx_hash)"
const DATABASE_CREATE_TRANSACTION_INDEX4 string = "CREATE INDEX IF NOT EXISTS " +
	"index_hash_idx ON transactions (tx_index,tx_hash,block_index)"

const DATABASE_CREATE_MESSAGE string = "CREATE TABLE IF NOT EXISTS message (" +
	"message_index INTEGER," +
	"block_index INTEGER," +
	"block_hash TEXT," +
	"block_time INTEGER," +
	"tx_index INTEGER," +
	"tx_hash TEXT," +
	"account TEXT," +
	"source TEXT," +
	"destination TEXT," +
	"message BLOB," +
	"FOREIGN KEY (block_index,block_hash,block_time) REFERENCES blocks(block_index,block_hash,block_time)," +
	"FOREIGN KEY (tx_index,tx_hash) REFERENCES transactions (tx_index,tx_hash)," +
	"PRIMARY KEY (message_index,account,source,destination))"
const DATABASE_CREATE_MESSAGE_INDEX string = "CREATE INDEX IF NOT EXISTS " +
	"message_index_idx ON message (message_index) "
const DATABASE_CREATE_MESSAGE_INDEX1 string = "CREATE INDEX IF NOT EXISTS " +
	"account_idx ON message (account)"
const DATABASE_CREATE_MESSAGE_INDEX2 string = "CREATE INDEX IF NOT EXISTS " +
	"source_idx ON message (source)"
const DATABASE_CREATE_MESSAGE_INDEX3 string = "CREATE INDEX IF NOT EXISTS " +
	"destination_idx ON message (destination)"

type DBTransaction struct {
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

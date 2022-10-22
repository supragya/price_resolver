package priceresolver

import (
	badger "github.com/dgraph-io/badger/v3"
	"github.com/dgraph-io/badger/v3/options"
)

type Engine struct {
	db     DB
	ethrpc EthRPC
	cex    CEX
}

func ConnectDB(path string) (*badger.DB, error) {
	opts := badger.DefaultOptions(path).
		WithCompression(options.None).
		WithSyncWrites(true)
	db, err := badger.Open(opts)
	return db, err
}

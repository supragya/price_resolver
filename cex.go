package priceresolver

import (
	"math/big"
	"time"
)

type CEX interface {
	GetPrice(string, time.Time) (*big.Float, error)
}

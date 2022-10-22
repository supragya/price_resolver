package priceresolver

import (
	"github.com/ethereum/go-ethereum/common"
)

type EthRPC interface {
	GetTokenName([]common.Address) (string, error)
	GetCLLatestAnswer(common.Address) error
	GetUniswapLatestPrice(common.Address, common.Address) error
}

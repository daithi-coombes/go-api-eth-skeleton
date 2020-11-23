package dao

import (
	"math/big"
	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/ethclient"
)

type DAO interface{
	GetBalance() big.Int
	GetProposal(big.Int) map[string]string
}

func (dao TECGarden) GetBalance() big.Int{
	var ret big.Int

	return ret
}

func (dao TECGarden) GetProposal() map[string]string{
	var ret map[string]string

	return ret
}

package dao

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// The DAO interface
type DAO interface {
	TotalProposals() (*big.Int, error)
}

// NewDAO Factory method for getting a doa instance for a specific protocol. Eg protocol=aragon
// TODO: finish this
func NewDAO(protocol string, organisation common.Address, endpoint string) (DAO, error) {

	var instance DAO

	client, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	contractAddr := common.HexToAddress("0x839c81ecdc41ff9a4fe291240a6961d482d19234")
	instance, err = NewConvictionBeta(contractAddr, client)
	if err != nil {
		return instance, err
	}

	return instance, nil
}

// TotalProposals Get total proposals for an Aragon organisation
// @uses ConvictionBeta
func (d ConvictionBeta) TotalProposals() (*big.Int, error) {
	var total *big.Int

	total, err := d.ProposalCounter(&bind.CallOpts{})
	if err != nil {
		return total, err
	}

	return total, nil
}

package dao

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// The DAO interface
type DAO interface {
	TotalProposals() *big.Int
	GetProposal(big.Int) interface{}
}

// NewDAO Factory method for getting a doa instance for a specific protocol. Eg protocol=aragon
// TODO: finish this
func NewDAO(protocol string, organisation common.Address) DAO {

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

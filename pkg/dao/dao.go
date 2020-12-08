package dao

import (
	"math/big"

	"github.com/daithi-coombes/go-api-eth-skeleton/pkg/dao/TECGardens"
	"github.com/ethereum/go-ethereum/common"
)

// The DAO interface
type DAO interface {
	TotalProposals() (*big.Int, error)           // total number of proposals
	GetOrganization() (interface{}, error)       // struct with organization details. Values will be interpolated with related template as message downstream
	GetProposal(uid string) (interface{}, error) // get details for a single proposal. Values will be interpolated with related template as message downstream
	GetProposals(limit int) (interface{}, error) // details for all proposals.
}

// NewDAO Factory method for getting a doa instance for a specific protocol. Eg protocol=aragon | TECGardens
// TODO: finish this
func NewDAO(protocol string, organisation common.Address, endpoint string) (DAO, error) {

	var instance DAO

	switch protocol {
	case "TECGardens":
		instance, err := TECGardens.NewTECGardens(organisation, endpoint)
		if err != nil {
			return instance, err
		}
	}

	return instance, nil
}

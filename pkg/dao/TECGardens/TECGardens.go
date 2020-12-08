package TECGardens

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TECGardens struct {
	binding *ConvictionBeta
}
type foo struct{}

func NewTECGardens(organization common.Address, endpoint string) (TECGardens, error) {

	var instance TECGardens

	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return instance, err
	}

	// contractAddr := common.HexToAddress("0x839c81ecdc41ff9a4fe291240a6961d482d19234")
	instance.binding, err = NewConvictionBeta(organization, client)
	if err != nil {
		return instance, err
	}

	return instance, nil
}

// TotalProposals Get total proposals for an Aragon organisation
// @uses ConvictionBeta
func (t TECGardens) TotalProposals() (*big.Int, error) {
	var total *big.Int

	total, err := t.binding.ProposalCounter(&bind.CallOpts{})
	if err != nil {
		return total, err
	}

	return total, nil
}

func (t TECGardens) GetOrganization() (interface{}, error) {
	return foo{}, nil
}
func (t TECGardens) GetProposal(uid string) (interface{}, error) {
	return foo{}, nil
}
func (t TECGardens) GetProposals(limit int) (interface{}, error) {
	var ret []foo
	return append(ret, foo{}), nil
	// return [foo{},], nil
}

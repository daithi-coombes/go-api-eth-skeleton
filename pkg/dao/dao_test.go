package dao

// Organization: 0xd0817aa0f770d024f42f0222dba37536a05118dc
// Conviction: 0x839c81ecdc41ff9a4fe291240a6961d482d19234

import (
	"math/big"
	"testing"

	"github.com/daithi-coombes/go-api-eth-skeleton/pkg/dao/TECGardens"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestTotalProposals(t *testing.T) {

	underTest, err := helperGetDAO()
	actual, err := underTest.TotalProposals()
	if err != nil {
		t.Error(err)
	}

	expected := big.NewInt(4)
	assert.Equal(t, expected, actual)
}

func TestGetProposals(t *testing.T) {
	underTest, err := helperGetDAO()
}

func TestGetOrganization(t *testing.T) {
	// underTest, err := helperGetDAO()
	// actual, err := underTest.GetOrganization()
	// if err != nil {
	// 	t.Error(err)
	// }

	// expected := big.NewInt(4)
	// assert.Equal(t, expected, actual)
}

func helperGetDAO() (DAO, error) {
	convictionAddress := "0x839c81ecdc41ff9a4fe291240a6961d482d19234"
	contractAddr := common.HexToAddress(convictionAddress)
	return TECGardens.NewTECGardens(contractAddr, "https://rpc.xdaichain.com")
}

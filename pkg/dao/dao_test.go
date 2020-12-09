package dao

// Organization: 0xd0817aa0f770d024f42f0222dba37536a05118dc
// Conviction: 0x839c81ecdc41ff9a4fe291240a6961d482d19234

import (
	"math/big"
	"reflect"
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
	if err != nil {
		t.Error(err)
	}

	proposals, err := underTest.GetProposals(100)
	if err != nil {
		t.Error(err)
	}
	actual := reflect.ValueOf(proposals)

	assert.Equal(t, 3, actual.Len())
}

func TestGetProposal(t *testing.T) {

	underTest, err := helperGetDAO()
	if err != nil {
		t.Error(err)
	}

	proposal, err := underTest.GetProposal("2")
	if err != nil {
		t.Error(err)
	}
	actual := reflect.ValueOf(proposal)

	assert.Equal(t, 1, actual.Len())
}

func TestGetOrganization(t *testing.T) {
	underTest, err := helperGetDAO()
	if err != nil {
		t.Error(err)
	}

	org, err := underTest.GetOrganization()
	if err != nil {
		t.Error(err)
	}
	response := org.(TECGardens.GraphQLResponse)
	actualApps := reflect.ValueOf(response.Data.Organization.Apps)

	assert.Equal(t, 13, actualApps.Len())
}

func helperGetDAO() (DAO, error) {
	convictionAddress := "0x839c81ecdc41ff9a4fe291240a6961d482d19234"
	contractAddr := common.HexToAddress(convictionAddress)
	return TECGardens.NewTECGardens(contractAddr, "https://rpc.xdaichain.com")
}

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
	actual1, err := underTest.TotalProposals()
	if err != nil {
		t.Error(err)
	}

	expected1 := big.NewInt(4)
	assert.Equal(t, expected1, actual1)

	type Data struct{ Total *big.Int }
	data := Data{big.NewInt(4)}
	actual2, err := underTest.ParseTemplate("TECGardens/total", data)
	if err != nil {
		t.Error(err)
	}
	expected2 := "There are 4 proposals"

	assert.Equal(t, expected2, actual2)
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
	actual1 := reflect.ValueOf(proposals)

	expected := 3
	assert.Equal(t, expected, actual1.Len())

	actual2, err := underTest.ParseTemplate("TECGardens/proposals", proposals)
	if err != nil {
		t.Error(err)
	}

	expected2 := "Proposals:\n\n\n - Abstain proposal\n\n - test\n\n - Donate to the Commons Stack\n"
	assert.Equal(t, expected2, actual2)
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

	expected := 1
	assert.Equal(t, expected, actual.Len())

	actual2, err := underTest.ParseTemplate("TECGardens/proposal", proposal)
	if err != nil {
		t.Error(err)
	}

	expected2 := "\n    test\n\n    Requesting:\n        100000000000\n    Total Stakes:\n        3\n    Forum Post:\n        https://forum.tecommons.org/t/translate-to-spanish-the-onboarding-and-resources-guide-doc/51/3\n"
	assert.Equal(t, expected2, actual2)
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

	expected := 13
	assert.Equal(t, expected, actualApps.Len())

	actual2, err := underTest.ParseTemplate("TECGardens/organization", response.Data.Organization)
	if err != nil {
		t.Error(err)
	}

	expected2 := "The following apps are installed:\n\n\n - dandelion-voting\n\n - tollgate\n\n - marketplace-presale\n\n - marketplace-controller\n\n - vault\n\n - conviction-beta\n\n - marketplace-bancor-market-maker\n\n - \n\n - vault\n\n - \n\n - redemptions\n\n - token-manager\n\n - \n"
	assert.Equal(t, expected2, actual2)
}

func helperGetDAO() (DAO, error) {
	convictionAddress := "0x839c81ecdc41ff9a4fe291240a6961d482d19234"
	contractAddr := common.HexToAddress(convictionAddress)
	return TECGardens.NewTECGardens(contractAddr, "https://rpc.xdaichain.com")
}

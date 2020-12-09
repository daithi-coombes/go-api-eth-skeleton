package TECGardens

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TECGardens The main struct for this package's instance
type TECGardens struct {
	binding *ConvictionBeta
}
type Stake struct {
	Amount string `json:"amount"`
	Entity string `json:"entity"`
}
type Proposal struct {
	Link            string  `json:"link"`
	Name            string  `json:"name"`
	Number          string  `json:"number"`
	RequestedAmount string  `json:"requestedAmount"`
	Stakes          []Stake `json:"stakes"`
}
type Organization struct {
	Apps []App `json:"apps"`
}
type App struct {
	Organization interface{} `json:"organization"`
	Repo         struct {
		Address     string `json:"address"`
		LastVersion struct {
			Artifact    string `json:"artifact"`
			CodeAddress string `json:"codeAddress"`
			ContentURI  string `json:"contentUri"`
			Manifest    string `json:"manifest"`
		} `json:"lastVersion"`
		Name     string      `json:"name"`
		Registry interface{} `json:"registry"`
	} `json:"repo"`
}
type GraphQLResponse struct {
	Data struct {
		Proposals    []Proposal   `json:"proposals",omitempty`
		Organization Organization `json:"organization",omitempty`
	} `json:"data"`
}

type foo struct{} // debug code

// TheGraphRequest thegraph.com request json for graphql http(s) queries
type TheGraphRequest struct {
	OperationName string `json:"operationName"`
	Query         string `json:"query"`
	Variables     struct {
		AppFilter  struct{} `json:"appFilter"`
		OrgAddress string   `json:"orgAddress"`
	} `json:"variables"`
	AppFilter  struct{} `json:"appFilter"`
	OrgAddress string   `json:"orgAddress"`
}

// NewTECGardens Factory function for getting a DAO instance for deployed TECommons Gardens Template
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

// GetOrganization Returns the on chain details for the organization querring thegraph.com
func (t TECGardens) GetOrganization() (interface{}, error) {

	var body GraphQLResponse
	payload := []byte("{\"query\": \"query Organization($orgAddress: String!, $appFilter: App_filter!, $first: Int) {  organization(id: $orgAddress) { apps(where: $appFilter, first: $first) {   ...App_app   __typename } __typename  }}fragment App_app on App {  address  appId  isForwarder  isUpgradeable  repoName  implementation { address __typename  }  organization { address __typename  }  version { ...Version_version __typename }  repo { ...Repo_repo __typename  }  roles { ...Role_role __typename  }}fragment Repo_repo on Repo {  address  name  node  registry { address __typename  }  lastVersion { ...Version_version __typename  }  versions { ...Version_version __typename  }}fragment Version_version on Version {  semanticVersion  codeAddress  contentUri  artifact  manifest}fragment Role_role on Role {  roleHash manager  appAddress  grantees { ...Permission_permission __typename  }}fragment Permission_permission on Permission {  appAddress  allowed  granteeAddress  roleHash  params { argumentId operationType argumentValue __typename  }}\",\"operationName\": \"Organization\",\"variables\": {\"appFilter\": {},\"orgAddress\": \"0xd0817aa0f770d024f42f0222dba37536a05118dc\"}}")
	url := "https://api.thegraph.com/subgraphs/name/1hive/aragon-xdai"

	resp, err := t.doGraphQL(payload, url)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()

	body, err = t.parseGraphQL(resp)
	if err != nil {
		return body, err
	}

	return body, nil
}

// GetProposal Get a single proposal
func (t TECGardens) GetProposal(number string) (interface{}, error) {

	var body GraphQLResponse
	payload := []byte("{\"query\": \"query Proposals($appAddress: String!, $first: Int!, $skip: Int!) {proposals(where: {appAddress: $appAddress,number: " + number + "}, first: $first, skip: $skip) {  id  number  name  link  creator  beneficiary  requestedAmount  status  totalTokensStaked  stakes(first: 1000) {entity amount __typename  }  appAddress  __typename}}\",\"operationName\": \"Proposals\",\"variables\": {\"appAddress\": \"0x839c81ecdc41ff9a4fe291240a6961d482d19234\",\"first\": 1000,\"skip\": 0}}")
	url := "https://api.thegraph.com/subgraphs/name/1hive/aragon-conviction-voting-xdai"

	resp, err := t.doGraphQL(payload, url)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()

	body, err = t.parseGraphQL(resp)
	if err != nil {
		return body, err
	}

	return body.Data.Proposals, nil
}

// GetProposals Querries thegraph.com for proposals
func (t TECGardens) GetProposals(limit int) (interface{}, error) {

	var body GraphQLResponse
	payload := []byte("{\"query\":\"query Proposals($appAddress: String!, $first: Int!, $skip: Int!) {proposals(where: {appAddress: $appAddress}, first: $first, skip: $skip) {  id  number  name  link  creator  beneficiary  requestedAmount  status  totalTokensStaked  stakes(first: 1000) {    entity    amount    __typename  }  appAddress  __typename}}\",\"operationName\":\"Proposals\",\"variables\":{\"appAddress\":\"0x839c81ecdc41ff9a4fe291240a6961d482d19234\",\"first\":1000,\"skip\":0}}")
	url := "https://api.thegraph.com/subgraphs/name/1hive/aragon-conviction-voting-xdai"

	resp, err := t.doGraphQL(payload, url)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()

	body, err = t.parseGraphQL(resp)
	if err != nil {
		return body, err
	}

	return body.Data.Proposals, nil
}

// doGraphQL Run http requests for thegraph.com GraphQL
func (t TECGardens) doGraphQL(payload []byte, url string) (*http.Response, error) {

	// TODO: don't hard code this, should be in config
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return &http.Response{}, err
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	return client.Do(request)
}

func (t TECGardens) parseGraphQL(resp *http.Response) (GraphQLResponse, error) {
	var body GraphQLResponse

	_bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	if resp.Status != "200 OK" {
		return body, errors.New(string(_bytes))
	}

	if err := json.Unmarshal(_bytes, &body); err != nil {
		return body, err
	}

	return body, nil
}

package dao

// Organization: 0xd0817aa0f770d024f42f0222dba37536a05118dc
// Conviction: 0x839c81ecdc41ff9a4fe291240a6961d482d19234

import (
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func TestTotalProposals(t *testing.T) {

	client, err := ethclient.Dial("https://rpc.xdaichain.com")
	if err != nil {
		log.Fatal(err)
	}

	contractAddr := common.HexToAddress("0x839c81ecdc41ff9a4fe291240a6961d482d19234")
	underTest, err := NewConvictionBeta(contractAddr, client)
	actual, err := underTest.TotalProposals()
	if err != nil {
		t.Error(err)
	}

	expected := big.NewInt(3)
	assert.Equal(t, expected, actual)
}

// Uncomment and run to get the full organisation details.
// @TODO get these details from the aragon/organisation registry
func TestGetOrganisationGraph(t *testing.T) {

	// type requestPayload struct {
	// 	OperationName string `json:"operationName"`
	// 	Query         string `json:"query"`
	// 	Variables     struct {
	// 		AppFilter  struct{} `json:"appFilter"`
	// 		OrgAddress string   `json:"orgAddress"`
	// 	} `json:"variables"`
	// 	AppFilter  struct{} `json:"appFilter"`
	// 	OrgAddress string   `json:"orgAddress"`
	// }

	// type Organization struct {
	// 	Data struct {
	// 		Organization struct {
	// 			Apps []struct {
	// 				Organization interface{} `json:"organization"`
	// 				Repo         struct {
	// 					Address     string `json:"address"`
	// 					LastVersion struct {
	// 						Artifact    string `json:"artifact"`
	// 						CodeAddress string `json:"codeAddress"`
	// 						ContentURI  string `json:"contentUri"`
	// 						Manifest    string `json:"manifest"`
	// 					} `json:"lastVersion"`
	// 					Name     string      `json:"name"`
	// 					Registry interface{} `json:"registry"`
	// 				} `json:"repo"`
	// 			} `json:"apps"`
	// 		} `json:"organization"`
	// 	} `json:"data"`
	// }

	// payload := []byte("{\"query\": \"query Organization($orgAddress: String!, $appFilter: App_filter!, $first: Int) {  organization(id: $orgAddress) { apps(where: $appFilter, first: $first) {   ...App_app   __typename } __typename  }}fragment App_app on App {  address  appId  isForwarder  isUpgradeable  repoName  implementation { address __typename  }  organization { address __typename  }  version { ...Version_version __typename }  repo { ...Repo_repo __typename  }  roles { ...Role_role __typename  }}fragment Repo_repo on Repo {  address  name  node  registry { address __typename  }  lastVersion { ...Version_version __typename  }  versions { ...Version_version __typename  }}fragment Version_version on Version {  semanticVersion  codeAddress  contentUri  artifact  manifest}fragment Role_role on Role {  roleHash manager  appAddress  grantees { ...Permission_permission __typename  }}fragment Permission_permission on Permission {  appAddress  allowed  granteeAddress  roleHash  params { argumentId operationType argumentValue __typename  }}\",\"operationName\": \"Organization\",\"variables\": {\"appFilter\": {},\"orgAddress\": \"0xd0817aa0f770d024f42f0222dba37536a05118dc\"}}")
	// url := "https://api.thegraph.com/subgraphs/name/1hive/aragon-xdai"
	// request, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	// if err != nil {
	// 	panic(err)
	// }

	// request.Header.Set("Content-Type", "application/json")
	// client := &http.Client{}
	// resp, err := client.Do(request)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	// _bytes, _ := ioutil.ReadAll(resp.Body)
	// // var body interface{}
	// var body Organization
	// if err := json.Unmarshal(_bytes, &body); err != nil {
	// 	spew.Dump(err)
	// 	panic(err)
	// }

	// spew.Dump(body)
	// fmt.Println("response Body:", string(body))
}

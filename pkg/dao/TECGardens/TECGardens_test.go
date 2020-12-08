package TECGardens

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	// // "github.com/davecgh/go-spew/spew"
)

type requestPayload struct {
	OperationName string `json:"operationName"`
	Query         string `json:"query"`
	Variables     struct {
		AppFilter  struct{} `json:"appFilter"`
		OrgAddress string   `json:"orgAddress"`
	} `json:"variables"`
	AppFilter  struct{} `json:"appFilter"`
	OrgAddress string   `json:"orgAddress"`
}

// Uncomment and run to get the full organisation details.
// @TODO get these details from the aragon/organisation Registry contract
func TestGetOrganisationGraph(t *testing.T) {

	type Organization struct {
		Data struct {
			Organization struct {
				Apps []struct {
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
				} `json:"apps"`
			} `json:"organization"`
		} `json:"data"`
	}

	payload := []byte("{\"query\": \"query Organization($orgAddress: String!, $appFilter: App_filter!, $first: Int) {  organization(id: $orgAddress) { apps(where: $appFilter, first: $first) {   ...App_app   __typename } __typename  }}fragment App_app on App {  address  appId  isForwarder  isUpgradeable  repoName  implementation { address __typename  }  organization { address __typename  }  version { ...Version_version __typename }  repo { ...Repo_repo __typename  }  roles { ...Role_role __typename  }}fragment Repo_repo on Repo {  address  name  node  registry { address __typename  }  lastVersion { ...Version_version __typename  }  versions { ...Version_version __typename  }}fragment Version_version on Version {  semanticVersion  codeAddress  contentUri  artifact  manifest}fragment Role_role on Role {  roleHash manager  appAddress  grantees { ...Permission_permission __typename  }}fragment Permission_permission on Permission {  appAddress  allowed  granteeAddress  roleHash  params { argumentId operationType argumentValue __typename  }}\",\"operationName\": \"Organization\",\"variables\": {\"appFilter\": {},\"orgAddress\": \"0xd0817aa0f770d024f42f0222dba37536a05118dc\"}}")
	url := "https://api.thegraph.com/subgraphs/name/1hive/aragon-xdai"
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	_bytes, _ := ioutil.ReadAll(resp.Body)
	// var body interface{}
	var body Organization
	if err := json.Unmarshal(_bytes, &body); err != nil {
		// spew.Dump(err)
		panic(err)
	}

	// spew.Dump(body)
	// fmt.Printf("response Body: %v\n", body)
}

func TestGetProposalsGraph(t *testing.T) {
	type stake struct {
		Amount string `json:"amount"`
		Entity string `json:"entity"`
	}
	type proposal struct {
		Link            string  `json:"link"`
		Name            string  `json:"name"`
		Number          string  `json:"number"`
		RequestedAmount string  `json:"requestedAmount"`
		Stakes          []stake `json:"stakes"`
	}
	type response struct {
		Data struct {
			Proposals []proposal `json:"proposals"`
		} `json:"data"`
	}

	// {"query":"query Proposals($appAddress: String!, $first: Int!, $skip: Int!) {\n  proposals(where: {appAddress: $appAddress}, first: $first, skip: $skip) {\n    id\n    number\n    name\n    link\n    creator\n    beneficiary\n    requestedAmount\n    status\n    totalTokensStaked\n    stakes(first: 1000) {\n      entity\n      amount\n      __typename\n    }\n    appAddress\n    __typename\n  }\n}\n","operationName":"Proposals","variables":{"appAddress":"0x839c81ecdc41ff9a4fe291240a6961d482d19234","first":1000,"skip":0}}
	payload := []byte("{\"query\":\"query Proposals($appAddress: String!, $first: Int!, $skip: Int!) {proposals(where: {appAddress: $appAddress}, first: $first, skip: $skip) {  id  number  name  link  creator  beneficiary  requestedAmount  status  totalTokensStaked  stakes(first: 1000) {    entity    amount    __typename  }  appAddress  __typename}}\",\"operationName\":\"Proposals\",\"variables\":{\"appAddress\":\"0x839c81ecdc41ff9a4fe291240a6961d482d19234\",\"first\":1000,\"skip\":0}}")

	url := "https://api.thegraph.com/subgraphs/name/1hive/aragon-conviction-voting-xdai"
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	_bytes, _ := ioutil.ReadAll(resp.Body)
	if resp.Status != "200 OK" {
		log.Fatal(string(_bytes))
		return
	}

	var body response
	if err := json.Unmarshal(_bytes, &body); err != nil {
		panic(err)
	}

	// spew.Dump(body)
}

# TECommons DAO Chat Bot

This bot will listen for requests from muliple social media channels and reply with results taken from blockchain and related DLT protocols.
ie: Its a bridge between your Discord, Slack, Telegram channels and the current state.

## Setting up channels

Here are quick instructions with links to releveant information on how to setup with different channels:
 - Telegram
    ```
    TELEGRAM_BOT_TOKEN=my_telegram_api
    go run main.go --organization 0xd0817aa0f770d024f42f0222dba37536a05118dc --conviction 0x839c81ecdc41ff9a4fe291240a6961d482d19234 --endpoint https://rpc.xdaichain.com
    ```

    Once your bot is registered you can get teh total number of proposals:
    ```
    Total Proposals 3
    ```
 -  Discord
    ...
 -  Slack
    ...


## Usage

In these examples we will be using the ConvictionVotingBeta smart contract deployed through the TECommons Gardens Template.
Post the following message inside a social media channel that is correctly setup (see above):
```bash
/tec dao get proposal $x
...
```

## Development

When you are adding a DAO from a new protocol to this package you will need to make sure you use the `dao.DAO` interface in `./pkg/dao.go`. For example to add the `Ocean` protocol you would create:
`./pkg/dao/Ocean/Ocean.go`
```golang
package OceanDAO

type Ocean struct{}

func NewOcean() (Ocean, error) { return Ocean{}, nil}

func (o Ocean) TotalProposals() (*big.Int, error)           // total number of proposals
func (o Ocean) GetOrganization() (interface{}, error)       // struct with organization details. Values will be interpolated with related template as message downstream
func (o Ocean) GetProposal(uid string) (interface{}, error) // get details for a single proposal. Values will be interpolated with related template as message downstream
func (o Ocean) GetProposals(limit int) (interface{}, error) // details for all proposals.
```

Next add your new protocol to the switch statement, currently at: `./pkg/dao.go::NewDAO()`
```golang
...
	switch protocol {
	case "TECGardens":
		instance, err := TECGardens.NewTECGardens(organisation, endpoint)
		if err != nil {
			return instance, err
        }
    case "Ocean":
		instance, err := OceanDAO.NewOcean(organisation, endpoint)
		if err != nil {
			return instance, err
        }
	}
...
```

To run the tests agains just your protocol then edit the helper method in `pkg/dao_test.go::helperGetDAO()` and run the tests:
```golang
func helperGetDAO() (DAO, error) {
	convictionAddress := "0x0"
	contractAddr := common.HexToAddress(convictionAddress)
	return OceanDAO.NewOcean(contractAddr, "https://rinkeby.infura.com")
}
```

... and run the tests:
```
make test
```

## Roadmap

 1. Enforce config file. Different DAO's will consume different params. This means a need to have structured and nested params of different complexities. Enforce YAML
 2. Factory methods (eg `TECGardens.NewTECGardens`) should take a config type to suit 1 above. Keep factory methods decoupled from all other packages.
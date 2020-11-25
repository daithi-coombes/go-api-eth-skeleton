Skeleton for creating API's on Ethereum based projects in golang

Usage:
```
go run main.go --help
```

```
TELEGRAM_BOT_TOKEN=my_telegram_api
go run main.go --organization 0xd0817aa0f770d024f42f0222dba37536a05118dc --conviction 0x839c81ecdc41ff9a4fe291240a6961d482d19234 --endpoint https://rpc.xdaichain.com
```

Once your bot is registered you can get teh total number of proposals:
```
Total Proposals 3
```

Requirements:
 - [x] TDD
 - [x] Simulated Backend
 - [x] Http Client
 - [ ] Websocket Client

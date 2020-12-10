package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/daithi-coombes/go-api-eth-skeleton/pkg/dao"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// Update The telegram api payload
type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

// Message The telegram api message field
type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

// Chat The telegram api chat field
type Chat struct {
	ID int `json:"id"`
}

var org, conv, endpoint string

// Handler Handles requests from the telegram group
func Handler(res http.ResponseWriter, req *http.Request) {
	var body Update
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		fmt.Printf("Error decoding request body: %s\n", err)
		return
	}

	// check for command
	words := strings.Split(body.Message.Text, " ")
	if words[0] != "/tec" {
		return
	}
	log.Printf("[request] %v\n", words)

	response := do(words[1], words[2], words[3])
	log.Printf("[response] %s\n", response)

	chatID := body.Message.Chat.ID
	var telegramAPI string = "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_BOT_TOKEN") + "/sendMessage"

	resp, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id": {strconv.Itoa(chatID)},
			"text":    {response},
		})
	if err != nil {
		fmt.Printf("Error posting to telegram api: %s\n", err)
		return
	}
	defer resp.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(resp.Body)
	if errRead != nil {
		fmt.Printf("Error reading telegram response with message: %s: %s", response, errRead)
		return
	}
	bodyString := string(bodyBytes)
	log.Printf("Telegram response for '%s': %s", response, bodyString)
}

// Run The main command to run the bot
func Run(cmd *cobra.Command) {
	var err error
	org, err = cmd.Flags().GetString("organization")
	if err != nil {
		log.Fatalf("Error parsing organization flag: %s\n", err)
	}
	conv, err = cmd.Flags().GetString("conviction")
	if err != nil {
		log.Fatalf("Error parsing organization flag: %s\n", err)
	}
	endpoint, err = cmd.Flags().GetString("endpoint")
	if err != nil {
		log.Fatalf("Error parsing endpoint flag: %s\n", err)
	}
	http.ListenAndServe(":3000", http.HandlerFunc(Handler))
}

func do(domain string, action string, model string) string {
	if domain == "dao" {

		dao, err := dao.NewDAO("aragon", common.HexToAddress(conv), endpoint)
		if err != nil {
			log.Fatalf("Uable to initiate DAO instance: %s\n", err)
		}

		switch model {
		case "proposal":
			if action == "total" {
				total, err := dao.TotalProposals()
				if err != nil {
					return fmt.Sprintf("Error getting total proposals: %s\n", err)
				}
				resp, err := dao.ParseTemplate("TECGardens/total", total)
				if err != nil {
					return fmt.Sprintf("Error parsing total proposals template: %s\n", err)
				}
				return resp
			}
			return "proposal this!"
		default:
			return "my dog eat the proposal!"
		}
	} else {
		return "even robots get confused... or were you only saying hi?"
	}
}

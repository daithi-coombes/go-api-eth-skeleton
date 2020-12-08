package client

import (
	"testing"
)

func TestWebSocket(t *testing.T) {

	// addr := "api.thegraph.com"

	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt)

	// u := url.URL{Scheme: "wss", Host: addr, Path: "/subgraphs/name/1hive/aragon-conviction-voting-xdai"}
	// log.Printf("Connecting to: %s\n", u.String())

	// c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	// if err != nil {
	// 	spew.Dump(u)
	// 	log.Fatalf("dial error: %s", err)
	// }
	// defer c.Close()

	// done := make(chan struct{})

	// go func() {
	// 	defer close(done)
	// 	for {
	// 		_, message, err := c.ReadMessage()
	// 		if err != nil {
	// 			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
	// 				log.Printf("error: %v", err)
	// 			}
	// 			log.Printf("Read Error: %s\n", err)
	// 			return
	// 		}
	// 		log.Printf("recv: %s", message)
	// 	}
	// }()

	// ticker := time.NewTicker(time.Second)
	// defer ticker.Stop()

	// for {
	// 	select {
	// 	case <-done:
	// 		return
	// 	case t := <-ticker.C:
	// 		if err := c.WriteMessage(websocket.TextMessage, []byte(t.String())); err != nil {
	// 			log.Printf("Write Error: %s\n", err)
	// 			return
	// 		}
	// 	case <-interrupt:
	// 		log.Printf("interrupt")

	// 		if err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
	// 			log.Printf("Close Error: %s\n", err)
	// 		}
	// 		select {
	// 		case <-done:
	// 		case <-time.After(time.Second):
	// 		}
	// 		return
	// 	}
	// }
}

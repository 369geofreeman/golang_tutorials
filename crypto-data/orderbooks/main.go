package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type Trade struct {
	EventType       string     `json:"e"`
	EventTime       int64      `json:"E"`
	TransactionTime int64      `json:"T"`
	Symbol          string     `json:"s"`
	U               int64      `json:"U"`
	Uu              int64      `json:"u"`
	Pu              int64      `json:"pu"`
	Bids            [][]string `json:"b"`
	Asks            [][]string `json:"a"`
}

func main() {

	var url string = "wss://fstream.binance.com/ws/bnbusdt@depth@500ms"
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// producer: read data stream from websocket and send to channel
	input := make(chan Trade)

	go func() {
		// read from the websocket
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				break
			}
			// unmarshal the message
			// fmt.Println(message['b'])
			var trade Trade
			json.Unmarshal(message, &trade)
			// send the trade to the channel
			input <- trade
		}
		close(input)
	}()

	fmt.Println(input)

	output := make(chan Trade)
	go func() {
		for trade := range input {
			output <- trade
		}
		close(output)
	}()

	for trade := range output {
		json, _ := json.Marshal(trade)
		fmt.Println(string(json))
	}

}

//  #################################################

// package main

// import (
// 	"encoding/json"
// 	"fmt"

// 	"github.com/gorilla/websocket"
// )

// type Trade struct {
// 	Exchange  string  `json:"exchange"`
// 	Base      string  `json:"base"`
// 	Quote     string  `json:"quote"`
// 	Direction string  `json:"direction"`
// 	Price     float64 `json:"price"`
// 	Volume    int64   `json:"volume"`
// 	Timestamp int64   `json:"timestamp"`
// 	PriceUsd  float64 `json:"priceUsd"`
// }

// func main() {

// 	c, _, err := websocket.DefaultDialer.Dial("wss://ws.coincap.io/trades/binance", nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer c.Close()

// 	// producer: read data stream from websocket and send to channel
// 	input := make(chan Trade)

// 	go func() {
// 		// read from the websocket
// 		for {
// 			_, message, err := c.ReadMessage()
// 			if err != nil {
// 				break
// 			}
// 			// unmarshal the message
// 			var trade Trade
// 			json.Unmarshal(message, &trade)
// 			// send the trade to the channel
// 			input <- trade
// 		}
// 		close(input)
// 	}()

// 	dogecoin := make(chan Trade)
// 	go func() {
// 		for trade := range input {
// 			if trade.Base == "dogecoin" && trade.Quote == "tether" {
// 				dogecoin <- trade
// 			}
// 		}
// 		close(dogecoin)
// 	}()

// 	// print the trades
// 	for trade := range dogecoin {
// 		json, _ := json.Marshal(trade)
// 		fmt.Println(string(json))
// 	}

// }

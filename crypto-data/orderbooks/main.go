package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
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

func readData(url string) {

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

	output := make(chan Trade)
	go func() {
		for trade := range input {
			output <- trade
		}
		close(output)
	}()

	for trade := range output {
		timestamp, _ := json.Marshal(trade.EventTime)
		bids, _ := json.Marshal(trade.Bids)
		asks, _ := json.Marshal(trade.Asks)

		//  push to database
		fmt.Println(string(timestamp))
		fmt.Println(string(bids))
		fmt.Println(string(asks))
	}
}

func writeToDatabase() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgress"
		password = "***"
		dbname   = "orderbook_data"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func main() {

	// var coin string = "bnbusdt"
	// url := fmt.Sprintf("wss://fstream.binance.com/ws/%v@10@500ms", coin)

	// readData(url)
	writeToDatabase()
}

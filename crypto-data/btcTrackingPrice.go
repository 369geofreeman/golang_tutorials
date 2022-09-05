package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Bina struct {
	Price string `json:"price"`
}

var btc Bina

func trackingPrice() {
	for {
		data, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT")
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(data.Body)
		if err != nil {
			log.Fatalln(err)
		}

		err = json.Unmarshal([]byte(body), &btc)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(btc.Price)
		time.Sleep(2 * time.Second)
		data.Body.Close()
	}
}

func main() {
	trackingPrice()
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"image/color"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"time"
// )

// type Bina struct {
// 	Price string `json:"price"`
// }

// var btc Bina

// func trackingPrice(c chan string) {
// 	for {
// 		data, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT")
// 		if err != nil {
// 			log.Fatalln(err)
// 		}

// 		body, err := ioutil.ReadAll(data.Body)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}

// 		err = json.Unmarshal([]byte(body), &btc)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}

// 		fmt.Println(btc.Price)
// 		c <- btc.Price
// 		time.Sleep(5 * time.Second)
// 		data.Body.Close()
// 	}
// }

// func main() {
// 	c := make(chan string)
// 	value := "WAITING"
// 	go trackingPrice(c)
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("BTC price")
// 	green := color.NRGBA{R: 11, G: 102, B: 35, A: 255}
// 	board := canvas.newText(value, green)
// 	go func() {
// 		for range time.Tick(time.Second) {
// 			value = <-c
// 			fmt.Println("Received:", value)
// 		}
// 	}()
// 	board.TextSize = 100
// 	board.TextStyle = fyne.TextStyle{Bold: true}
// 	myWindow.SetContent(board)
// 	myWindow.ShowAndRun()
// }

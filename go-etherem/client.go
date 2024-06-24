package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	//client, err := ethclient.Dial("http://localhost:8545")
	//if err != nil {
	//	log.Fatal(err)
	//}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections

}

package main

import (
	"log"
	"net/http"
	"os"
)

const difficulty = 1

type Block struct {
	Index      int
	Timestamp  string
	Data       int
	Hash       string
	Prevhash   string
	Difficulty int
	Nonce      string
}

var Blockchain []Block

func main() {
	err := os.Getenv(".evn")
	if err != "" {
		log.Fatal(err)
	}
}

func run() error {
	return nil
}

func makeMuxRouter() http.Handler {
	HandlerFunc("/", handleGetBlockchain).Method("GET")
	HandlerFunc("/", handleWriteBlock).Method("POST")
}

func handleGetBlockchain() {}

func handleWriteBlock() {}

func respondWithJSON() {}

func isBlockValid() bool {
	return true
}

func calculateHash() string {
	return ""
}

func generateBlock() {}

func isHashValid() bool {
	return true
}

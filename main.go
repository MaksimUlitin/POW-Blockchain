package main

import (
	"net/http"
)

type Block struct{}

var Blockchain []Block

func main() {

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

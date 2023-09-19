package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
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

var (
	Blockchain []Block
	mutex      = &sync.Mutex{}
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := Block{}
		genesisBlock = Block{0, t.String(), calculateHash(genesisBlock), "", difficulty, ""}
		spew.Dump(genesisBlock)
		mutex.Lock()
		Blockchain = append(Blockchain, genesisBlock)
		mutex.Unlock()
	}()
	log.Fatal(run())
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

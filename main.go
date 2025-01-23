package main

import (
	"log"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/maksimUlitin/internal/blockchain"
	"github.com/maksimUlitin/internal/server"
)

const difficulty = 1

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := blockchain.Block{
			Index:      0,
			Timestamp:  t.String(),
			Data:       0,
			Hash:       blockchain.CalculateHash(blockchain.Block{}),
			Prevhash:   "",
			Difficulty: difficulty,
			Nonce:      "",
		}
		spew.Dump(genesisBlock)
		blockchain.Mutex.Lock()
		blockchain.Blockchain = append(blockchain.Blockchain, genesisBlock)
		blockchain.Mutex.Unlock()
	}()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

	log.Fatal(server.Run(port))
}

package main

import (
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
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
	mux := makeMuxRouter()
	portHttp := os.Getenv("PORT")
	log.Println("HTTP server is running and listening on port:", portHttp)
	s := &http.Server{
		Addr:           ":" + portHttp,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandlerFunc("/", handleGetBlockchain).Method("GET")
	muxRouter.HandlerFunc("/", handleWriteBlock).Method("POST")
	return muxRouter
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

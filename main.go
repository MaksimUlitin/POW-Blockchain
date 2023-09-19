package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/text/message"
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

type Message struct {
	Data int
}
var (
	Blockchain []Block
	mutex      = &sync.Mutex{}
	m = Message
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

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Blockchain, "", " ")
	if err != nil {
		
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type","application/json")

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w,r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	mutex.Lock()
	newBlock := generateBlock(Blockchain[len(Blockchain)-1], m.Data)
	mutex.Unlock()
	if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]){
		Blockchain = append(Blockchain, newBlock)
		spew.Dump(Blockchain)
	}
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	w.Header().Set("content-Type", "application/json")
	respons, err := json.MarshalIndent(payload, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: internal server error"))
	}
	
	w.WriteHeader(code)
	w.Write(respons)
	
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index +1
}

func calculateHash(w http.ResponseWriter, r *http.Request) string {
	return ""
}

func generateBlock(oldBlock Block, Data int) Block {}

func isHashValid(w http.ResponseWriter, r *http.Request) bool {
	return true
}

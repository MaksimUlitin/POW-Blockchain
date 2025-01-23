package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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

type Message struct {
	Data int
}

var (
	Blockchain []Block
	mutex      = &sync.Mutex{}
	m          Message
	newBlock   Block
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := Block{}
		genesisBlock = Block{0, t.String(), 0, calculateHash(genesisBlock), "", difficulty, ""}
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
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
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
	w.Header().Set("content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	mutex.Lock()
	newBlock := generateBlock(Blockchain[len(Blockchain)-1], m.Data)
	mutex.Unlock()
	if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
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
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.Prevhash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return false
}

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.Data) + block.Prevhash + block.Nonce

	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, Data int) Block {

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = Data
	newBlock.Prevhash = oldBlock.Hash
	newBlock.Difficulty = difficulty

	for i := 1; ; i++ {
		hex := fmt.Sprintf("%x", i)
		newBlock.Nonce = hex

		if !isHashValid(calculateHash(newBlock), newBlock.Difficulty) {
			fmt.Println(calculateHash(newBlock), "do more work")
			time.Sleep(time.Second)
			continue
		} else {
			fmt.Println(calculateHash(newBlock), "work done")
			newBlock.Hash = calculateHash(newBlock)
			break
		}

	}
	return newBlock
}

func isHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", 1)
	return strings.HasPrefix(hash, prefix)
}

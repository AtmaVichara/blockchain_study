package main

import (
  "fmt"
  "crypto/sha256"
  "encoding/json"
  "encoding/hex"
  "io"
  "log"
  "net/http"
  "time"
  "os"

  "github.com/davecgh/go-spew/spew"
  "github.com/joho/godotenv"
  "github.com/gorilla/mux"
)

type Block struct {
  Index     int
  Timestamp string
  BPM       int
  Hash      string
  PrevHash  string
}

type Message struct {
  BPM int
}

var Blockchain []Block

func calculateHash(block Block) string {
  record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
  hashed := sha256.Sum256([]byte(record))
  return hex.EncodeToString(hashed[:])
}

func generateBlock(oldBlock Block, BPM int) (Block, error) {
  t := time.Now()

  newBlock := &Block{
    Index: oldBlock.Index + 1,
    Timestamp: t.String(),
    BPM: BPM,
    PrevHash: oldBlock.Hash,
  }
  newBlock.Hash = calculateHash(*newBlock)

  return *newBlock, nil
}

func isBlockValid(newBlock, oldBlock Block) bool {
  switch {
    case newBlock.Index != oldBlock.Index + 1:
      return false
    case newBlock.PrevHash != oldBlock.Hash:
      return false
    case newBlock.Hash != calculateHash(newBlock):
      return false
    default:
      return true
  }
}

func replaceChain(newBlocks []Block) {
  if len(newBlocks) > len(Blockchain) {
    Blockchain = newBlocks
  }
}

func run() error {
  mux := makeRouter()
  httpAdder := os.Getenv("ADDR")
  fmt.Println("Listening on :" + httpAdder)
  timeout := 10 * time.Second

  s := &http.Server{
    Addr: ":" + httpAdder,
    Handler: mux,
    ReadTimeout: timeout,
    WriteTimeout: timeout,
    MaxHeaderBytes: 1 << 20,
  }

  if err := s.ListenAndServe(); err != nil {
    return err
  }

  return nil
}

func makeRouter() http.Handler {
  router := mux.NewRouter()
  router.HandleFunc("/", handleGetBlockchain).Methods("GET")
  router.HandleFunc("/", handleWriteBlockchain).Methods("POST")
  return router
}

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
  bytes, err := json.MarshalIndent(Blockchain, "", " ")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  io.WriteString(w, string(bytes))
}

func handleWriteBlockchain(w http.ResponseWriter, r *http.Request) {
  var m Message

  decoder := json.NewDecoder(r.Body)
  if err := decoder.Decode(&m); err != nil {
    respondWithJSON(w, r, http.StatusBadRequest, r.Body)
    return
  }
  defer r.Body.Close()

  newBlock, err := generateBlock(Blockchain[len(Blockchain)-1], m.BPM)
  if err != nil {
    respondWithJSON(w, r, http.StatusInternalServerError, m)
    return
  }

  if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
    newBlockChain := append(Blockchain, newBlock)
    replaceChain(newBlockChain)
    spew.Dump(Blockchain)
  }

  respondWithJSON(w, r, http.StatusCreated, newBlock)
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, status int, payload interface{}) {
  response, err := json.MarshalIndent(payload, "", " ")
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("HTTP 500: Internal Server Error"))
    return
  }

  w.WriteHeader(status)
  w.Write([]byte(response))
}

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal(err.Error())
  }

  go func() {
    t := time.Now()
    genesisBlock := Block{0, t.String(), 0, "", ""}
    spew.Dump(genesisBlock)
    Blockchain = append(Blockchain, genesisBlock)
  }()
  log.Fatal(run())
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params,omitempty"`
	ID      int           `json:"id"`
}

type RPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   interface{} `json:"error"`
	ID      int         `json:"id"`
}

type BlockNumberRequest struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	ID      int    `json:"id"`
}

type BlockByNumberRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

type BlockchainClient struct {
	client *http.Client
	host   string
}

func NewBlockchainClient(host string) *BlockchainClient {
	return &BlockchainClient{
		client: &http.Client{},
		host:   host,
	}
}

func (bc *BlockchainClient) makeRPCRequest(request RPCRequest) (*RPCResponse, error) {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}
	resp, err := bc.client.Post(bc.host, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}
	var rpcResponse RPCResponse
	err = json.Unmarshal(body, &rpcResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}
	return &rpcResponse, nil
}

func (bc *BlockchainClient) GetBlockNumber(w http.ResponseWriter, r *http.Request) {
	var req BlockNumberRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if req.JSONRPC != "2.0" || req.Method != "eth_blockNumber" || req.ID != 2 {
		http.Error(w, "Request must match specified JSON-RPC format", http.StatusBadRequest)
		return
	}
	rpcReq := RPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_blockNumber",
		ID:      2,
	}
	response, err := bc.makeRPCRequest(rpcReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (bc *BlockchainClient) GetBlockByNumber(w http.ResponseWriter, r *http.Request) {
	var req BlockByNumberRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if req.JSONRPC != "2.0" || req.Method != "eth_getBlockByNumber" || req.ID != 2 || len(req.Params) != 2 {
		http.Error(w, "Request must match specified JSON-RPC format", http.StatusBadRequest)
		return
	}
	rpcReq := RPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  req.Params,
		ID:      2,
	}
	response, err := bc.makeRPCRequest(rpcReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	client := NewBlockchainClient("https://polygon-rpc.com/")
	router := mux.NewRouter()
	router.HandleFunc("/block/number", client.GetBlockNumber).Methods("POST")
	router.HandleFunc("/block/by-number", client.GetBlockByNumber).Methods("POST")
	log.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

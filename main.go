package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Block struct {
	Index         int    `json:"index"`
	Timestamp     int    `json:"timestamp"`
	Data          string `json:"String"`
	Previous_hash string `json:"previous_hash"`
	Hash          string `json:"hash"`
}

type searchID struct {
	Index int `json:"index"`
}

var Blockchain []Block
var nextID int = 1

/**
 * Function to create a new Blocks.
 *
 * @param data
 * 			http.ResponseWriter - To respond to the server
 *			http.Request - get request from the server
 * @return
 *           None
 */
func createBlock(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request", http.StatusMethodNotAllowed)
		return
	}
	var eachBlock Block
	// Read body of the POST request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invlaid Request Data", http.StatusBadRequest)
		return
	}
	eachBlock.Index = nextID
	nextID++
	// Parsing the JSON data
	err = json.Unmarshal(body, &eachBlock)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	Blockchain = append(Blockchain, eachBlock)
	//Response back with thesame message
	w.Header().Set("Content/Type", "application/json")
	json.NewEncoder(w).Encode(Blockchain)
}

/**
 * Function to get all Block from the blockchain.
 *
 * @param data
 * 			http.ResponseWriter - To respond to the server
 *			http.Request - get request from the server
 * @return
 *           None
 */
func getBlock(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content/Type", "application/json")
	json.NewEncoder(w).Encode(Blockchain)
}

/**
 * Function to get Blocks by Index from the blockchain.
 *
 * @param data
 * 			http.ResponseWriter - To respond to the server
 *			http.Request - get request from the server
 * @return
 *           None
 */
func getBlockByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invlaid Request Data", http.StatusBadRequest)
		return
	}
	var SearchByID searchID
	// Parsing the JSON data
	err = json.Unmarshal(body, &SearchByID)
	fmt.Println(SearchByID)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	var blockFound *Block
	for i := 0; i < len(Blockchain); i++ {
		if Blockchain[i].Index == SearchByID.Index {
			blockFound = &Blockchain[i]
			break
		}
	}
	w.Header().Set("Content/Type", "application/json")
	json.NewEncoder(w).Encode(blockFound)
}

/**
 * Function to update the Block in the blockchain.
 *
 * @param data
 * 			http.ResponseWriter - To respond to the server
 *			http.Request - get request from the server
 * @return
 *           None
 */
func updateBlockByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invlaid Request Data", http.StatusBadRequest)
		return
	}
	var SearchByID Block
	// Parsing the JSON data
	err = json.Unmarshal(body, &SearchByID)
	fmt.Println(SearchByID)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	for i := 0; i < len(Blockchain); i++ {
		if Blockchain[i].Index == SearchByID.Index {
			Blockchain[i].Timestamp = SearchByID.Timestamp
			Blockchain[i].Data = SearchByID.Data
			Blockchain[i].Previous_hash = SearchByID.Previous_hash
			Blockchain[i].Hash = SearchByID.Hash
			break
		}
	}
	w.Header().Set("Content/Type", "application/json")
	json.NewEncoder(w).Encode(Blockchain)
}

/**
 * Function to delete the Block from the blockchain.
 *
 * @param data
 * 			http.ResponseWriter - To respond to the server
 *			http.Request - get request from the server
 * @return
 *           None
 */
func deleteBlockByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invlaid Request Data", http.StatusBadRequest)
		return
	}
	var SearchByID searchID
	// Parsing the JSON data
	err = json.Unmarshal(body, &SearchByID)
	fmt.Println(SearchByID)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	var newBlockChain []Block
	for i := 0; i < len(Blockchain); i++ {

		// ChatGPT response to delete the object but I dodn't understand properly.
		// if Blockchain[i].Index == SearchByID.Index {
		// 	// Remove the person by shifting elements to the left
		// 	Blockchain = append(Blockchain[:i], Blockchain[i+1:]...)
		// 	i--
		// }

		// This is not the proper way of removing object from array.
		if Blockchain[i].Index != SearchByID.Index {
			newBlockChain = append(newBlockChain, Blockchain[i])
		}
	}
	Blockchain = newBlockChain
	w.Header().Set("Content/Type", "application/json")
	json.NewEncoder(w).Encode(Blockchain)
}

func main() {
	http.HandleFunc("/createBlock", createBlock)
	http.HandleFunc("/getBlock", getBlock)
	http.HandleFunc("/getBlockByID", getBlockByID)
	http.HandleFunc("/updateBlockByID", updateBlockByID)
	http.HandleFunc("/deleteBlockByID", deleteBlockByID)
	fmt.Println("Server running on port: 4455")
	log.Fatal(http.ListenAndServe(":4455", nil))
}

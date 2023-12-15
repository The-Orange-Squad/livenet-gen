package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// ErrorResponse is a struct for representing error responses
type ErrorResponse struct {
	Error string `json:"error"`
}

// TokenResponse is a struct for representing token responses
type TokenResponse struct {
	Token string `json:"token"`
}

// Token is a struct that holds a 128-symbol alphanumeric string
type Token struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

// Tokens is a map of user IDs to tokens
var Tokens map[int]Token

// Symbols is a string of possible characters for the token
const Symbols = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// GenerateToken returns a random token of 128 symbols
func GenerateToken() Token {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 128)
	for i := range b {
		b[i] = Symbols[rand.Intn(len(Symbols))]
	}
	return Token{Value: string(b)}
}

// LoadTokens reads the livenet_t.json file and loads the tokens into the map
func LoadTokens() {
	data, err := ioutil.ReadFile("livenet_t.json")
	if err != nil {
		log.Fatal(err)
	}

	var tokensMap map[string]Token  // Change the type to map[string]Token
	err = json.Unmarshal(data, &tokensMap)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the map with string keys to a map with integer keys
	intTokens := make(map[int]Token)
	for key, token := range tokensMap {
		id, err := strconv.Atoi(key)
		if err != nil {
			log.Fatal(err)
		}
		intTokens[id] = token
	}

	Tokens = intTokens
}


// SaveTokens writes the tokens map to the livenet_t.json file
func SaveTokens() {
	data, err := json.Marshal(Tokens)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("livenet_t.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
// GetHandler handles the GET requests and tries to retrieve a token for the given user ID
func GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	token, ok := Tokens[id]
	if !ok {
		http.Error(w, "No token found for this user ID", http.StatusNotFound)
		return
	}

	response := TokenResponse{Token: token.Value}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// SetHandler handles the SET requests and generates a new token for the given user ID
func SetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	_, ok := Tokens[id]
	if ok {
		http.Error(w, "A token already exists for this user ID", http.StatusConflict)
		return
	}
	token := GenerateToken()
	Tokens[id] = token
	SaveTokens()

	response := TokenResponse{Token: token.Value}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RewriteHandler handles the REWRITE requests and re-generates a token for the given user ID
func RewriteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	_, ok := Tokens[id]
	if !ok {
		http.Error(w, "No token found for this user ID", http.StatusNotFound)
		return
	}
	token := GenerateToken()
	Tokens[id] = token
	SaveTokens()

	response := TokenResponse{Token: token.Value}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteHandler handles the DELETE requests and deletes the token for the given user ID
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	_, ok := Tokens[id]
	if !ok {
		http.Error(w, "No token found for this user ID", http.StatusNotFound)
		return
	}

	// Delete the token for the given user ID
	delete(Tokens, id)
	SaveTokens()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Token deleted successfully"})
}



func main() {
	// Load the tokens from the file
	LoadTokens()

	// Create a new router
	r := mux.NewRouter()

	// Register the routes and handlers
	r.HandleFunc("/get/{id}", GetHandler).Methods("GET")
	r.HandleFunc("/set/{id}", SetHandler).Methods("GET")        // Change "SET" to "POST"
	r.HandleFunc("/rewrite/{id}", RewriteHandler).Methods("GET") // Change "REWRITE" to "POST"
	r.HandleFunc("/delete/{id}", DeleteHandler).Methods("GET")   // Change "DELETE" to "POST"

	// Start the server
	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TokenRequest struct {
	Username string `json:"username"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintln(w, "Method not supported")
		return
	}

	// Read request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error reading request body: %v", err)
		return
	}

	var req TokenRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding JSON request body: %v", err)
		return
	}

	// check if we already have the token cached
	if token := validToken(req); token != "" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, toString(TokenResponse{Token: token}))
		return
	}

	// generate fresh token
	token, err := generateToken(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error decoding JSON request body: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, toString(TokenResponse{Token: token}))
}

func main() {
	http.HandleFunc("/token", handlerFunc)
	fmt.Println("Server starting on port 4040")
	err := http.ListenAndServe(":4040", nil)
	if err != nil {
		panic(err)
	}
}

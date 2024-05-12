package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var catalogue = map[string][]string{
	"john":   {"The Intelligent Investor", "The Immortals of Meluha"},
	"mary":   {"Atomic Habbits", "Train to Pakistan"},
	"daniel": {"India After Independence"},
}

type CatalogueResponse struct {
	Books []string `json:"books"`
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	// Extract username from path variable
	username := mux.Vars(r)["username"]
	books, ok := catalogue[username]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User: %s not found", username)
		return
	}

	if err := validateAccess(r.Header.Get("Authorization"), username); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorized access: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, toString(CatalogueResponse{Books: books}))
}

func main() {
	r := mux.NewRouter() // Use a mux for routing
	r.HandleFunc("/users/{username}/books", booksHandler)
	fmt.Println("Server starting on port 5050")
	err := http.ListenAndServe(":5050", r)
	if err != nil {
		panic(err)
	}
}

func validateAccess(token, username string) error {
	if len(token) == 0 {
		return fmt.Errorf("token not passed")
	}

	// To keep it simple, we are not verifying the token signature.
	parsedToken, _, err := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		return fmt.Errorf("error parsing the token: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("claims in jwt token is not map claims")
	}

	now := time.Now().Unix()
	if !claims.VerifyExpiresAt(now, true) {
		return fmt.Errorf("token is expired")
	}

	user, ok := claims["username"].(string)
	if !ok {
		return fmt.Errorf("invalid user in token")
	}

	if user != username {
		return fmt.Errorf("username in token: %s and username in path: %s don't match", user, username)
	}

	return nil
}

func toString(i interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}

package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func validToken(req TokenRequest) string {
	token := cache.get(req)
	if token == "" {
		return ""
	}

	parsedToken, _, err := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		return ""
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return ""
	}

	now := time.Now().Unix()
	if !claims.VerifyExpiresAt(now, true) {
		return ""
	}

	return token
}

// Mock the token generation call.
// Here, this could be some CLI command or API call to external service based on the use case.
func generateToken(req TokenRequest) (string, error) {
	// add 2s delay to mock remote request
	time.Sleep(2 * time.Second)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"exp":      time.Now().Add(2 * time.Minute).Unix(),
	})

	jwtString, err := token.SignedString([]byte("my-secret"))
	if err != nil {
		return "", fmt.Errorf("unable to encode jwt to string: %w", err)
	}

	cache.set(req, jwtString)
	return jwtString, nil
}

func toString(i interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}

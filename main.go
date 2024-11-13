package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/yaml.v2"
	"net/http"
)

func main() {
	fmt.Println("Starting vulnerable Go application...")

	// Using jwt-go to create a token (vulnerable version)
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println("Error creating token:", err)
	}
	fmt.Println("Generated Token:", tokenString)

	// Using yaml.v2 which has known vulnerabilities
	data := `
name: Vulnerable Application
version: 1.0.0
`
	var config map[string]interface{}
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		fmt.Println("Error parsing YAML:", err)
	}
	fmt.Println("Parsed YAML:", config)

	// Simple HTTP server without proper security headers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

const PORT_ENV = "FORTUNE_PORT"

func fortune() (string, error) {
	output, err := exec.Command("fortune").Output()
	if err != nil {
		return "", errors.New("Command `fortune` not found!")
	}
	return string(output), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	response, err := fortune()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(response))
}

func main() {
	port := os.Getenv(PORT_ENV)
	if port == "" {
		message := fmt.Sprintf("Missing environmental variable: %s", PORT_ENV)
		fmt.Println(message)
		return
	}
	http.HandleFunc("/", handler)
	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}

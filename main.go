package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func calculateSumHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var numbers []int64

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &numbers)
	if err != nil {
		WriteJSONResponse(w, http.StatusInternalServerError, "Invalid JSON body")
		return
	}

	var sum int64
	for _, number := range numbers {
		sum += number
	}

	jsonResponse, err := json.Marshal(sum)
	if err != nil {
		WriteJSONResponse(w, http.StatusInternalServerError, "Error getting JSON response")
		return
	}

	WriteJSONResponse(w, http.StatusOK, jsonResponse)
}

func main() {
	http.HandleFunc("/calculate", calculateSumHandler)

	// Start the HTTP server
	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func get_method_example() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	sbody := string(bytes)
	log.Printf("Response: %s", sbody)
}

func post_method_example() {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "John Doe",
		"email": "john.doe@example.com",
	})

	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)
	if err != nil {
		log.Fatalf("Error making POST request: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	sbody := string(body)
	log.Printf("Response: %s", sbody)
}

func main() {
	go get_method_example()
	go post_method_example()

	time.Sleep(5 * time.Second) // Wait for the goroutine to finish
}

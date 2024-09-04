package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	go func() {
		for {
			// Create a slice with a random length between 0 and 199
			ids := make([]int, rng.Int31n(200))

			for i := range ids {
				ids[i] = rng.Intn(99) + 1
			}

			// Example: print the generated ids
			fmt.Println(ids)

			// Prepare the structure for the JSON payload
			data := map[string][]int{
				"object_ids": ids,
			}

			// Convert the structure to JSON
			jsonData, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Error marshalling JSON:", err)
				return
			}

			// Create a new POST request
			url := "http://localhost:9090/callback"
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
			if err != nil {
				fmt.Println("Error creating request:", err)
				return
			}

			// Set headers if needed
			req.Header.Set("Content-Type", "application/json")

			// Send the request using http.Client
			client := &http.Client{
				Transport: &http.Transport{DisableKeepAlives: true},
			}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Error sending request:", err)
				continue
			}
			defer resp.Body.Close()

			// Read the response
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading response body:", err)
				continue
			}

			fmt.Println("Response status:", resp.Status)
			fmt.Println("Response body:", string(body))
			time.Sleep(5 * time.Second)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	fmt.Println("closing")

}

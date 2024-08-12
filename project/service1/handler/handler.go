package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project/service1/model"
)

func CallbackHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestContent model.CallbackRequest
	if err := json.NewDecoder(request.Body).Decode(&requestContent); err != nil {
		http.Error(writer, fmt.Sprintf("invalid request: %s", err), http.StatusBadRequest)
		return
	}

	if len(requestContent.ObjectIDs) == 0 {
		http.Error(writer, "no object IDs provided", http.StatusBadRequest)
		return
	}

	// Process the callback data (for now, just print it)
	log.Printf("Received callback data: %s\n", requestContent)

	_, err := writer.Write([]byte("ok"))
	if err != nil {
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

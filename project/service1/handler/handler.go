package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"project/service1/model"
	"project/service1/repository"
	"strconv"
	"time"
)

type LastSeenCallback struct {
	ObjectRepoInterface repository.ObjectRepoInterface
}

func (h *LastSeenCallback) CallbackHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Step 1: Read the raw request body
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, fmt.Sprintf("could not read request body: %s", err), http.StatusInternalServerError)
		return
	}

	// Step 2: Log the raw request body
	log.Printf("Received raw request: %s\n", string(body))

	var requestContent model.CallbackRequest
	if err := json.Unmarshal(body, &requestContent); err != nil {
		http.Error(writer, fmt.Sprintf("invalid request: %s", err), http.StatusBadRequest)
		return
	}

	if len(requestContent.ObjectIDs) == 0 {
		http.Error(writer, "no object IDs provided", http.StatusBadRequest)
		return
	}

	// Process the callback data (for now, just print it)
	log.Printf("Received callback data: %v\n", requestContent)

	ctx := context.Background()

	var callBack model.DBObject

	for i := range requestContent.ObjectIDs {
		callBack.LastSeen = time.Now()
		callBack.ID = requestContent.ObjectIDs[i]

		err = h.ObjectRepoInterface.Upsert(ctx, &callBack)

		if err != nil {
			log.Printf("Failed to insert record: %v", err)
			return
		}
	}

	_, err = writer.Write([]byte(strconv.Itoa(len(requestContent.ObjectIDs))))
	if err != nil {
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

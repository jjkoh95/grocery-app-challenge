package groceryservice

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"os"
	"sync"

	"cloud.google.com/go/firestore"
)

const (
	firebaseCollectionPrefix = "GroceryApp"
)

var (
	projectID = os.Getenv("PROJECTID")

	firestoreClient     *firestore.Client
	firestoreClientOnce sync.Once
)

// Grocery is the type for grocery data
type Grocery struct {
	UPC12Barcode       big.Int  `json:"UPC12Barcode"`       // this is the `PrimaryKey`
	UPC12BarcodeString string   `json:"UPC12BarcodeString"` // this is to support writing to database
	Brand              string   `json:"brand"`
	ProductName        string   `json:"productName"`
	KeyWords           []string `json:"keyWords"`
}

type searchPayload struct {
	SearchQuery string `json:"searchQuery"`
}

// GetGrocery is the endpoint to get all grocery data.
// If data is not found in cache, query from database
func GetGrocery(w http.ResponseWriter, r *http.Request) {
	// CORS handler
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var search searchPayload
	err := json.NewDecoder(r.Body).Decode(&search)
	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	firestoreClientOnce.Do(initDB())

	groceries, err := getGroceryFromDB(context.Background(), search)
	if err != nil {
		http.Error(w, "Failed to query from DB", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(groceries)
	if err != nil {
		http.Error(w, "Error handling query response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// UpsertGrocery is the endpoint to upsert grocery data.
// This should upsert both database and cache.
func UpsertGrocery(w http.ResponseWriter, r *http.Request) {
	// CORS handler
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var grocery Grocery
	err := json.NewDecoder(r.Body).Decode(&grocery)
	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	firestoreClientOnce.Do(initDB())

	GetKeyWords(&grocery)
	err = upsertGroceryToDB(context.Background(), &grocery)
	if err != nil {
		http.Error(w, "Failed to upsert data", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully upserted data"))
}

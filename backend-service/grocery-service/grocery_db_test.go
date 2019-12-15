package groceryservice_test

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"

	groceryservice "github.com/jjkoh95/grocery-app-challenge/backend-service/grocery-service"
)

func TestGetKeyWords(t *testing.T) {

	testGroceryPayload := []byte(`{"UPC12Barcode":111122223333,"brand":"Famous Brand","productName":"Useful Product"}`)

	var testGrocery groceryservice.Grocery
	err := json.Unmarshal(testGroceryPayload, &testGrocery)
	if err != nil {
		t.Error("Expected to unmarshal payload successfully")
	}

	groceryservice.GetKeyWords(&testGrocery)

	expectedKeyWords := []string{
		"use", "useful", "product", "famous", "prod", "",
	}

	// making it fancy
	var wg sync.WaitGroup
	for _, expected := range expectedKeyWords {
		wg.Add(1)
		go func(wg *sync.WaitGroup, expected string) {
			for _, kw := range testGrocery.KeyWords {
				if kw == expected {
					wg.Done()
					return
				}
			}
			t.Error(fmt.Sprintf(`"%s" not found!`, expected))
			wg.Done()
		}(&wg, expected)
	}
	wg.Wait()
}

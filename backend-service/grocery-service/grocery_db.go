package groceryservice

import (
	"context"
	"log"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

func initDB() func() {
	return func() {
		ctx := context.Background()
		config := &firebase.Config{
			ProjectID: projectID,
		}
		app, err := firebase.NewApp(ctx, config)
		if err != nil {
			log.Printf("Firebase %v", err)
		}
		firestoreClient, err = app.Firestore(ctx)
		if err != nil {
			log.Printf("Firestore %v", err)
		}
	}
}

func getGroceryFromDB(ctx context.Context, search searchPayload) ([]Grocery, error) {
	groceryRef := firestoreClient.Collection(firebaseCollectionPrefix)
	query := groceryRef.Where("keyWords", "array-contains", search.SearchQuery)
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		log.Printf("Firestore %v", err)
		return nil, err
	}

	groceries := make([]Grocery, len(docs))
	for i, doc := range docs {
		doc.DataTo(&groceries[i])
		groceries[i].KeyWords = nil
		groceries[i].UPC12Barcode.SetString(groceries[i].UPC12BarcodeString, 10) // have to cast to big int
	}
	return groceries, nil
}

func upsertGroceryToDB(ctx context.Context, grocery *Grocery) error {
	// mergeAll only support map data
	// TODO: How to fix big.Int
	groceryMap := map[string]interface{}{
		"UPC12BarcodeString": grocery.UPC12Barcode.String(),
		"brand":              grocery.Brand,
		"productName":        grocery.ProductName,
		"keyWords":           grocery.KeyWords,
	}

	_, err := firestoreClient.Collection(firebaseCollectionPrefix).Doc(grocery.UPC12Barcode.String()).Set(
		ctx,
		groceryMap,
		firestore.MergeAll,
	)
	if err != nil {
		log.Printf("Firestore %v", err)
		return err
	}
	return nil
}

// GetKeyWords is to populate keywords for search
func GetKeyWords(grocery *Grocery) {
	productNames := strings.Split(grocery.ProductName, " ")
	brands := strings.Split(grocery.Brand, " ")

	keyWordsMap := make(map[string]interface{})
	for _, p := range append(productNames, brands...) {
		if len(p) < 3 {
			continue
		}
		for i := 0; i <= len(p); i++ {
			keyWordsMap[strings.ToLower(p[0:i])] = nil
		}
	}

	allKeys := make([]string, len(keyWordsMap))
	for k := range keyWordsMap {
		allKeys = append(allKeys, k)
	}
	grocery.KeyWords = allKeys
}

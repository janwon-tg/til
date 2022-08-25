package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func store(ctx context.Context, client *firestore.Client) {
	// add collection
	_, _, err := client.Collection("samples2").Add(ctx, map[string]interface{}{
		"first":  "Tom",
		"second": "じろう",
		"born":   2017,
	})

	if err != nil {
		log.Fatalf("Failed adding alovelace :%v", err)
	}

}

func addDocWithID(ctx context.Context, client *firestore.Client) {
	var data = make(map[string]interface{})

	_, err := client.Collection("cities").Doc("new-city-id").Set(ctx, data)
	if err != nil {
		log.Printf("An error has occured: %s", err)
	}
}

func addDocWithoutID(ctx context.Context, client *firestore.Client) {
	_, _, err := client.Collection("cities").Add(ctx, map[string]interface{}{
		"name":    "Tokyo",
		"country": "Japan",
	})
	if err != nil {
		log.Printf("An error has occured: %s", err)
	}
}

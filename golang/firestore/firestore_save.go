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

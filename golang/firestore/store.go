package main

import (
	"context"
	"log"
	"time"

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

func addDocWithID(id string, ctx context.Context, client *firestore.Client) {
	var data = make(map[string]interface{})

	_, err := client.Collection("cities").Doc(id).Set(ctx, data)
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

func addDocAsMap(ctx context.Context, client *firestore.Client) error {
	_, err := client.Collection("cities").Doc("LA").Set(ctx, map[string]interface{}{
		"name":    "Los Angels",
		"state":   "CA",
		"country": "USA",
	})

	if err != nil {
		log.Printf("An error has occurred = %+v\n", err)
	}

	return err
}

func addDocDataTypes(ctx context.Context, client *firestore.Client) error {
	doc := make(map[string]interface{})
	doc["stringExample"] = "Hello world!"
	doc["booleanExample"] = true
	doc["numberExample"] = 3.141592
	doc["dateExample"] = time.Now()
	doc["arrayExample"] = []interface{}{5, true, "hello"}
	doc["nullExample"] = nil
	doc["objectExample"] = map[string]interface{}{
		"a": 5,
		"b": true,
	}

	_, err := client.Collection("data").Doc("one").Set(ctx, doc)
	if err != nil {
		log.Printf("An error has occurred = %+v\n", err)
	}

	return err
}

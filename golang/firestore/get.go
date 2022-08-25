package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func getCollections(ctx context.Context, client *firestore.Client) error {
	iter := client.Collections(ctx)
	for {
		collRef, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return err
		}
		fmt.Printf("Found collection with id: %s\n", collRef.ID)
	}
	return nil
}

func getAllDocs(ctx context.Context, client *firestore.Client) error {
	iter := client.Collection("samples2").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return err
		}
		fmt.Println(doc.Data())
	}
	return nil
}

func search(ctx context.Context, client *firestore.Client) error {
	iter := client.Collection("samples2").Where("born", ">=", 2010).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(doc.Data())
	}
	return nil
}

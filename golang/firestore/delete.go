package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func deleteDoc(ctx context.Context, client *firestore.Client) error {

	_, err :=client.Collection("cities").Doc("DC").Delete(ctx)
	if err != nil {
		log.Printf("err = %+v\n", err)
	}

	return err
}

func deleteField(ctx context.Context, client *firestore.Client) error {
	_, err := client.Collection("cities").Doc("BJ").Update(ctx, []firestore.Update{
		{
			Path: "capital",
			Value: firestore.Delete,
		},
	})
	if err != nil {
		log.Printf("err = %+v\n", err)
	}
	return err
}

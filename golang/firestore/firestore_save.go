package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("/path/to/credentials_file")
	log.Print("NewApp")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	log.Print("Firestore")
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// add collection
	_, _, err = client.Collection("samples2").Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})

	if err != nil {
		log.Fatalf("Failed adding alovelace :%v", err)
	}

	log.Print("Close")
	defer client.Close()
}

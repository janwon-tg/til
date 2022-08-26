package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	log.Print("Start")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed read .evn : %v", err)
	}

	p := os.Getenv("CREDENTIALS_FILE")
	sa := option.WithCredentialsFile(p)

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("Failed NewApp: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln("Failed new Firestore: %v", err)
	}

	addDocWithID("DC", ctx, client)
	addDocWithID("BJ", ctx, client)
	addDocWithoutID(ctx, client)
	addDocAsMap(ctx, client)
	addDocDataTypes(ctx, client)

	log.Print("Call Store")
	store(ctx, client)

	log.Print("Call getCollections")
	getCollections(ctx, client)

	log.Print("Call getAllDocs")
	getAllDocs(ctx, client)

	log.Print("Call search")
	search(ctx, client)

	log.Print("Delete Doc")
	deleteDoc(ctx, client)
	log.Print("Delete Field")
	deleteField(ctx, client)

	log.Print("Close")
	defer client.Close()
}

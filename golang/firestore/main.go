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

	log.Print("Call Store")
	store(ctx, client)

	log.Print("Close")
	defer client.Close()
}

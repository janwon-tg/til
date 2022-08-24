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

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Failed read .env")
	}

	p := os.Getenv("CREDENTIALS_FILE")
	log.Print(p)
	sa := option.WithCredentialsFile(p)
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
		"born":  1816,
	})

	if err != nil {
		log.Fatalf("Failed adding alovelace :%v", err)
	}

	log.Print("Close")
	defer client.Close()
}

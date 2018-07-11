package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var idToken = "dummy"

func main() {
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
		return
	}

	client, err := app.Auth()
	if err != nil {
		log.Fatalln(err)
	}
	token, err := client.VerifyIDToken(idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	log.Printf("Verified ID token: %v\n", token)
}
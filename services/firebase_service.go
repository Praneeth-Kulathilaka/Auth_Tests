package services

import (
	"log"
	"context"
  
	firebase "firebase.google.com/go"
	// "firebase.google.com/go/auth"
  
	"google.golang.org/api/option"
  )

var FirebaseApp *firebase.App

func InitFirebase() {
	opt := option.WithCredentialsFile("firebase-adminsdk.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase: %v", err)
	}
	FirebaseApp = app
}


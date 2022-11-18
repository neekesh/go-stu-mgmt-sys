package infrastructure

import (
	"context"
	"fmt"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var AuthClient *auth.Client

func SetupFirebase() {
	serviceAccountKeyFilePath, err := filepath.Abs("./FirebaseServiceKey.json")
	fmt.Println(serviceAccountKeyFilePath)
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error(Fire base app )")
	}

	//Firebase Auth
	auth, err := app.Auth(context.Background())
	if err != nil {
		fmt.Println("firebase auththeicaiton error")
		panic("Firebase Authentication  error(Client Creation)")
	}

	AuthClient = auth

}

package infrastructure

import (
	"context"
	"fmt"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func NewFirebaseApp() *firebase.App {
	fmt.Println("called firebase app")
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

	return app
}

func NewFirebaseAuth(app *firebase.App) *auth.Client {
	fmt.Println("this is awesome authentication")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//Firebase Auth
	auth, err := app.Auth(ctx)
	if err != nil {
		fmt.Println("firebase auththeicaiton error")
		panic("Firebase Authentication  error(Client Creation)")
	}
	return auth

}

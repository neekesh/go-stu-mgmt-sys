package services

import (
	"context"

	"firebase.google.com/go/auth"
)

type FirebaseService struct {
	fbAuth *auth.Client
}

func NewFirebaseService(fbClient *auth.Client) FirebaseService {
	return FirebaseService{
		fbAuth: fbClient,
	}
}

func (fb *FirebaseService) VerifyToken(idToken string) (*auth.Token, error) {
	token, err := fb.fbAuth.VerifyIDToken(context.Background(), idToken)
	return token, err
}

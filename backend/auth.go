package main

import (
	"context"
	"errors"
	"net/http"
	"time"

	db "tourbackend/database/gen"
)

var COOKIE_LIFETIME = time.Hour * 24 * 7

func createHttpCookie(tokenValue string) *http.Cookie {
	return &http.Cookie{
		Name:     "auth_token",
		Value:    tokenValue,
		Expires:  time.Now().Add(COOKIE_LIFETIME),
		HttpOnly: true,
		Secure:   IS_DEPLOYED,
	}
}

func validateToken(token string, queries *db.Queries, ctx context.Context) (*db.User, error) {

	authInfo, err := queries.GetUserBySessionToken(ctx, token)
	if err != nil {
		return nil, err
	}

	if authInfo.ExpiresAt <= time.Now().Unix() {
		return nil, errors.New("session expired")
	}

	return &db.User{
		ID:        authInfo.UserID,
		FirstName: authInfo.FirstName,
		LastName:  authInfo.LastName,
		Hash:      authInfo.Hash,
		Email:     authInfo.Email,
	}, nil
}

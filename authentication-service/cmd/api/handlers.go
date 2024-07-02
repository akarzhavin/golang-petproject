package main

import (
	"authentification-service/models"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

var signingKey = []byte("secret")

type AuthResponseData struct {
	AccessToken  string `json:"access"`
	RefreshToken string `json:"refresh"`
}

type RefreshTokenRequest struct {
	Refresh string `json:"refresh"`
}

func (app *Config) RefreshToken(w http.ResponseWriter, r *http.Request) {
	requestPayload := RefreshTokenRequest{}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
	}

	if requestPayload.Refresh == "" {
		app.errorJSON(w, errors.New("Refresh token is empty"), http.StatusUnauthorized)
		return
	}

	token, err := app.Models.RefreshToken.GetOne(requestPayload.Refresh)
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	user, err := app.Models.User.GetOne(token.UserID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.sendNewTokens(w, r, user)
}

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	app.sendNewTokens(w, r, user)
}

func (app *Config) sendNewTokens(w http.ResponseWriter, r *http.Request, user *models.User) {
	accessToken, err := generateAccessToken(user)
	if err != nil {
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := generateRefreshToken(user)
	if err != nil {
		http.Error(w, "Failed to generate refresh token", http.StatusInternalServerError)
		return
	}

	err = refreshToken.Store()
	if err != nil {
		http.Error(w, "Failed to store refresh token to db", http.StatusInternalServerError)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in successfully"),
		Data: AuthResponseData{
			AccessToken:  accessToken,
			RefreshToken: refreshToken.Token,
		},
	}

	_ = app.writeJSON(w, http.StatusAccepted, payload)
}

func generateAccessToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   fmt.Sprintf("%v", user.ID),
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})
	return token.SignedString(signingKey)
}

func generateRefreshToken(user *models.User) (models.RefreshToken, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return models.RefreshToken{}, err
	}

	return models.RefreshToken{
		Token:     base64.URLEncoding.EncodeToString(b),
		UserID:    user.ID,
		Active:    true,
		UsedCount: 0,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}, nil
}

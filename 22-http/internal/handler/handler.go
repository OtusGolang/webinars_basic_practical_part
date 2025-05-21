package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"gohttp2/internal/config"
	"io"
	"net/http"
	"net/mail"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Base(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok", "data": "Main page"}`))
	return
}

func Login(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var user User
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"success": false}`))
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"success": false}`))
		return
	}
	_, err = mail.ParseAddress(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"success": false}`))
		return
	}

	userID, err := uuid.NewUUID()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"success": false}`))
	}

	secretKey := []byte(config.Cfg.Secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    userID.String(),
		"email": user.Username,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"success": true, "token": "%s"}`, tokenString)))
	return
}

func UserUpdate(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok", "data": "Main page"}`))
	return
}

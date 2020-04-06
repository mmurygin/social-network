package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
	"time"
)

var sessionSecret []byte

const cookieName = "X-Session-Token"

func init() {
	strSecret := os.Getenv("SESSION_SECRET")

	if strSecret == "" {
		log.Panic(errors.New("sessionSecret is not provided"))
	}

	sessionSecret = []byte(strSecret)
}

type Token struct {
	UserId int `json:"userId"`
	jwt.StandardClaims
}

func StoreSession(w http.ResponseWriter, userId int) error {
	log.Println("Store session")
	expiresAt := time.Now().Add(time.Hour)

	claims := Token{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(sessionSecret)

	if err != nil {
		log.Println(err)
		return err
	}

	cookie := http.Cookie{
		Name:    cookieName,
		Value:   tokenString,
		Expires: expiresAt,
	}

	http.SetCookie(w, &cookie)

	return nil
}

func GetSession(r *http.Request) (int, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return 0, err
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return sessionSecret, nil
	})

	if claims, ok := token.Claims.(*Token); ok && token.Valid {
		return claims.UserId, nil
	} else {
		return 0, err
	}
}

func CleanSession(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:    cookieName,
		Value:   "",
		MaxAge:  0,
		Expires: time.Now(),
	}

	http.SetCookie(w, &cookie)
}

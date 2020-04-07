package auth

import (
	"context"
	"github.com/mmurygin/social-network/data"
	"log"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := GetSession(r)

		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}

		user, err := data.QueryUser(userId)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}

		newContext := context.WithValue(r.Context(), "user", user)
		r = r.WithContext(newContext)

		next.ServeHTTP(w, r)
	})
}

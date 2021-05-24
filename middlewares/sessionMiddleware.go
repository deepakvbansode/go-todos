package middlewares

import (
	"context"
	"fmt"
	"go-todos/literals"
	"go-todos/models"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/microcosm-cc/bluemonday"
)

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		p := bluemonday.UGCPolicy()
		authToken := p.Sanitize(r.Header.Get("Authorization"))
		if authToken == "" {
			next.ServeHTTP(w, r)
			return
		}
		chunks := strings.SplitN(authToken, " ", 2)
		if len(chunks) == 2 {
			authToken = chunks[1]
		}

		jwtToken, err := jwt.Parse(authToken, nil)
		fmt.Println(err)
		if jwtToken == nil {
			next.ServeHTTP(w, r)
			return
		}

		claims, claimsOk := jwtToken.Claims.(jwt.MapClaims)
		if !claimsOk {
			next.ServeHTTP(w, r)
			return
		}
		session := &models.Session{}
		if sub, subOk := claims["sub"].(string); subOk {
			session.UserID = sub
		}
		if starID, starIDOk := claims["starId"].(string); starIDOk {
			session.StarID = starID
		}
		if email, emailOk := claims["email"].(string); emailOk {
			session.EmailID = email
		}
		ctx = context.WithValue(ctx, literals.RequestSessionUserKey, session)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

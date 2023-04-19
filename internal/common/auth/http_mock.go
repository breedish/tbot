package auth

import (
	"context"
	"net/http"

	"github.com/breedish/tbot/internal/common/server/httperr"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

func HttpMockMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var claims jwt.MapClaims
		token, err := request.ParseFromRequest(
			r,
			request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (i interface{}, e error) {
				return []byte("mock_secret"), nil
			},
			request.WithClaims(&claims),
		)
		if err != nil {
			httperr.BadRequest("unable-to-get-jwt", err, w, r)
			return
		}

		if !token.Valid {
			httperr.BadRequest("invalid-jwt", nil, w, r)
			return
		}

		ctx := context.WithValue(
			r.Context(),
			userContextKey,
			User{
				AccountId: claims["user_uuid"].(string),
				Username:  claims["name"].(string),
				Roles:     claims["roles"].([]string),
			},
		)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

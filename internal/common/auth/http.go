package auth

import (
	"context"
	"net/http"
	"strings"

	commonerrors "github.com/breedish/tbot/internal/common/errors"
)

type EverymeetAuthMiddleware struct {
	//AuthClient *auth.Client
}

func (a EverymeetAuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (a EverymeetAuthMiddleware) tokenFromHeader(r *http.Request) string {
	headerValue := r.Header.Get("Authorization")
	if len(headerValue) > 7 && strings.ToLower(headerValue[0:6]) == "bearer" {
		return headerValue[7:]
	}
	return ""
}

type User struct {
	AccountId string
	Username  string
	Roles     []string
}

type ctxKey int

const (
	userContextKey ctxKey = iota
)

var (
	NoUserInContextError = commonerrors.NewAuthorizationError("no user in context", "no-user-found")
)

func UserFromCtx(ctx context.Context) (User, error) {
	u, ok := ctx.Value(userContextKey).(User)
	if ok {
		return u, nil
	}

	return User{}, NoUserInContextError
}

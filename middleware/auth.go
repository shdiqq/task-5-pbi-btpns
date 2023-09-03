package middleware

import (
	"context"
	"net/http"

	"github.com/shdiqq/task-5-pbi-btpns-Shadiq/helpers"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		accessToken := request.Header.Get("Authorization")

		if accessToken == "" {
			helpers.Response(response, 401, "unauthorized", nil)
			return
		}

		user, err := helpers.ValidateToken(accessToken)
		if err != nil {
			helpers.Response(response, 401, err.Error(), nil)
			return
		}

		ctx := context.WithValue(request.Context(), "userinfo", user)
		next.ServeHTTP(response, request.WithContext(ctx))
	})
}

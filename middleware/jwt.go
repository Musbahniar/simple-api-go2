package middleware

import (
	"net/http"
	c "simple-api-go2/api/constants"
	"simple-api-go2/api/utils"
	"simple-api-go2/handler"
	"strconv"
	"strings"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		jwtToken := r.Header.Get("Authorization")
		if jwtToken == "" {
			handler.HttpError(w, http.StatusUnauthorized, c.AUTH_TOKEN_MISSING, r.URL)
			return
		}
		jwtToken = strings.Replace(jwtToken, "Bearer ", "", 1)
		claims, err := utils.VerifyJWTToken(jwtToken)
		if err != nil {
			handler.HttpError(w, http.StatusUnauthorized, c.INVALID_AUTH_TOKEN, r.URL)
			return
		}
		userId := strconv.Itoa(int(claims.Id))
		r.Header.Set("userName", claims.Name)
		r.Header.Set("userId", userId)
		next.ServeHTTP(w, r)
	})
}

package middleware

import (
	"net/http"
	c "simple-api-go2/api/constants"
	"simple-api-go2/api/utils"
	"simple-api-go2/handler"
)

const (
	ApiKeyHeader = "ApiKey"
)

func ApiKeyMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get(ApiKeyHeader)
		if apiKey == "" {
			handler.HttpError(w, http.StatusForbidden, c.UNAUTHRIZED_REQUEST, r.URL)
			return
		}
		_, err := utils.ValidateEnvVar(ApiKeyHeader, apiKey)
		if err != nil {
			handler.HttpError(w, http.StatusForbidden, c.UNAUTHRIZED_REQUEST, r.URL)
			return
		}
		next.ServeHTTP(w, r)
	})
}

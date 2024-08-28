package auth

import (
	"github.com/karthik-code78/ecom/shared/configure"
	"github.com/karthik-code78/ecom/shared/logging"
	"github.com/karthik-code78/ecom/shared/utils/http_utils"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

var token string

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		tokenHeader := req.Header.Get("Authorization")
		logging.Log.Info(tokenHeader)
		if tokenHeader == "" {
			http_utils.SendErrorResponse(res, "Auth token is missing, Please check!", http.StatusBadRequest)
			return
		}

		tokenStr := strings.Split(tokenHeader, " ")[1]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(configure.GetJwtSecretKey()), nil
		})

		if err != nil || !token.Valid {
			http_utils.SendErrorResponse(res, "Auth token is invalid, Please check!", http.StatusForbidden)
			return
		}

		next.ServeHTTP(res, req)
	})
}

func SetTokenForInternalCommunication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the token in the Authorization header
		token := "your-jwt-token-here" // You might want to retrieve this from a secure location
		r.Header.Set("Authorization", "Bearer "+token)

		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}

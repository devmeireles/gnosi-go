package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/devmeireles/gnosi-api/app/utils"
)

// AuthJwtVerify checks if the header contains a token
func AuthJwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("Authorization")
		header = strings.TrimSpace(header)

		if !utils.ValidateToken(header) {
			utils.ResErr(w, errors.New("unauthorized"), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

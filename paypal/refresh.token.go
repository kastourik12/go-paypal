package paypal

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RefreshToken(client CustomClient) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			err := client.RefreshToken()
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write([]byte(err.Error()))
				return
			}
			handler.ServeHTTP(rw, r)
		})
	}
}

package middleware

import (
	"net/http"

	"github.com/Jhairofc/gotwitter/routers"
)

//ValidateJWT nos permite validar la informacion que se obtiene del token
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		_, _, _, error := routers.ProcessToken(req.Header.Get("Authorization"))
		if error != nil {
			http.Error(res, "Error al obtener el token"+error.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(res, req)
	}
}

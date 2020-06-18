package middleware

import (
	"net/http"

	"github.com/Jhairofc/gotwitter/db"
)

//Checkdb middleware para conocer el estado de la db
func Checkdb(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if db.CheckConnectionDB() == 0 {
			http.Error(res, "Ha ocurrido un error, Conexión con la Base de Datos perdida", 500)
			return
		}
		//Si todo esta bien se pasa res y req a Routers para que continue
		next.ServeHTTP(res, req)
	}
}

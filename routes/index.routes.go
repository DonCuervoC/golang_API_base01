package routes

import "net/http"

// w = writer (como respondemos al cliente)
// r = reader (tener acceso a los parametros)
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World! listo para unas margaritas"))
}

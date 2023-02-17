package middlewares

import (
	"fmt"
	"net/http"
)

func (m *Middlewares) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Hi from Logger middleware")
		fmt.Println(req.URL, " : ", req.RemoteAddr)
		next.ServeHTTP(w, req)
	})
}

package logger

import (
	"log"
	"net/http"
)

func Logger(function http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		function.ServeHTTP(writer, request)
		log.Printf("%s\t%s\t%s\n", request.Host, request.URL.Path, request.URL.RawQuery)
	})
}

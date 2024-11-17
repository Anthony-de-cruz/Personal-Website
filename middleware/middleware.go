package middleware

import (
	"log"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

// Returns a response writer with a statusCode field.
func (wr *wrappedWriter) WriteHeader(statusCode int) {
	wr.ResponseWriter.WriteHeader(statusCode)
	wr.statusCode = statusCode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: wr,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapped, req)

		log.Println(wrapped.statusCode, req.Method, req.URL.Path, time.Since(start))
	})
}

type Middleware func(http.Handler) http.Handler

func CreateStack(funcs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(funcs) - 1; i >= 0; i-- {
			function := funcs[i]
			next = function(next)
		}
		return next
	}
}

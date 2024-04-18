package middleware

import (
	"net/http"

	"github.com/ak1m1tsu/go-libs/http/contextlib"
)

// ContextOpts is options for Context middleware configuration.
type ContextOpts struct {
	ServiceName string
}

// Context is middleware that wrap request id, request time, service name,
// and zerolog logger in context.
func Context(opts ContextOpts) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = contextlib.WithRequestID(ctx, r)
			ctx = contextlib.WithRequestTime(ctx)
			ctx = contextlib.WithServiceName(ctx, opts.ServiceName)
			ctx = contextlib.WithZeroLog(ctx)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// RequestID is middleware that adds header X-Request-Id to response headers.
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Request-Id", contextlib.GetRequestID(r.Context()))

		next.ServeHTTP(w, r)
	})
}

// ServiceName is middleware that adds header X-Service to response headers.
func ServiceName(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Service", contextlib.GetServiceName(r.Context()))

		next.ServeHTTP(w, r)
	})
}

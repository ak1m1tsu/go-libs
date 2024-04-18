package contextlib

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type ctxKey int

const (
	requestID ctxKey = iota
	requestTime
	serviceName
)

// WithRequestID returns context with request ID stored in headers as X-Request-Id
func WithRequestID(parent context.Context, r *http.Request) context.Context {
	reqID := r.Header.Get("X-Request-Id")
	if len(reqID) != 0 {
		parent = context.WithValue(parent, reqID, reqID)
	}
	return parent
}

// GetRequestID returns request ID as a string from context.
// If context doesn't have request id returns empty string.
func GetRequestID(ctx context.Context) string {
	if val, ok := ctx.Value(requestID).(string); ok {
		return val
	}

	return ""
}

// WithServiceName returns context with provided service name
func WithServiceName(parent context.Context, name string) context.Context {
	return context.WithValue(parent, serviceName, name)
}

// GetServiceName returns service name as a string from context.
// If context doesn't have service name returns empty string.
func GetServiceName(ctx context.Context) string {
	if val, ok := ctx.Value(serviceName).(string); ok {
		return val
	}

	return ""
}

// WithRequestTime returns context with request time (time.Now())
func WithRequestTime(parent context.Context) context.Context {
	return context.WithValue(parent, requestTime, time.Now())
}

// GetRequestTime returns request time as time.Time from context.
// If context doesn't have request time returns time.Now().
func GetRequestTime(ctx context.Context) time.Time {
	if val, ok := ctx.Value(requestTime).(time.Time); ok {
		return val
	}

	return time.Now()
}

// WithZeroLog returns context with zerolog logger
func WithZeroLog(parent context.Context) context.Context {
	return zerolog.New(os.Stdout).With().Caller().Timestamp().Logger().WithContext(parent)
}

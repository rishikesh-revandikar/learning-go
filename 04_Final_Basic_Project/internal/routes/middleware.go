package routes

import (
	"log/slog"
	"net/http"
	"time"
)

// LoggerMiddleware records the duration and details of every request
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Process the actual request
		next.ServeHTTP(w, r)
		
		// Log the results after completion
		slog.Info("Request Processed",
			"method", r.Method,
			"path", r.URL.Path,
			"duration", time.Since(start).String(),
		)
	})
}
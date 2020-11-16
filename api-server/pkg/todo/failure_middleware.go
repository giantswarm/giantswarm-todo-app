package todo

import (
	"math/rand"
	"net/http"
	"time"
)

// FailureMiddleware is a simple middleware to simulate HTTP errors
func FailureMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		num := rand.Int() % 10
		// return artificial error
		if num == 0 {
			http.Error(w, "This failed for no reason", http.StatusInternalServerError)
			return
		}
		if num == 1 {
			// delay the response
			time.Sleep(100 * time.Millisecond)
		}
		next.ServeHTTP(w, r)
	})
}

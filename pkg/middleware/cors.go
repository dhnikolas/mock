package middleware

import (
	"net/http"
)

// Мидлвер, который отлавливает панику и не дает процессу сервиса упасть с ошибкой.
// Этот мидлвайр должен использоватся для каждого роута
func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		next.ServeHTTP(w, r)
	})
}

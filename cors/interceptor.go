package cors

import (
	"net/http"
	"strings"

	"github.com/Eldius/cors-interceptor-go/config"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if origin != "" {
			if config.IsOriginAllowed(origin) {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", strings.Join(config.GetCORSAllowedMethods(), ", "))
				w.Header().Set("Access-Control-Allow-Headers", strings.Join(config.GetCORSAllowedHeaders(), ", "))
				if r.Method == http.MethodOptions {
					w.WriteHeader(200)
				}
			}
		}

		if r.Method != http.MethodOptions {
			next.ServeHTTP(w, r)
		}
	})
}

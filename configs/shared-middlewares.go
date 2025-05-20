package configs

import "net/http"

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") 
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		next.ServeHTTP(w, r)
	})
}


func EnableWebSecurityDefault(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'") // Restringe fontes de scripts
		w.Header().Set("X-Frame-Options", "DENY")                       // Evita iframes não autorizados
		w.Header().Set("X-Content-Type-Options", "nosniff")             // Previne MIME Sniffing
		w.Header().Set("X-XSS-Protection", "1; mode=block")             // Ativa proteção contra XSS
		//w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload") // Força HTTPS
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin") // Controle de Referer
		next.ServeHTTP(w, r)
	})
}


func LoadMiddlewares(next http.Handler) http.Handler{
	middlewareChain := EnableCORS(EnableWebSecurityDefault(next))
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		middlewareChain.ServeHTTP(w, r)
	})
}
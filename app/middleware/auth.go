package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) != 2 {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}
		tokenString := splitToken[1]
		fmt.Println("Received token :", tokenString)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})
		if err != nil {
			fmt.Println("Error parsing token:", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			r.Header.Set("user_id", fmt.Sprintf("%v", claims["user_id"]))
			next.ServeHTTP(w, r)
		} else {
			fmt.Println("Token is not valid")
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
	})
}

func JWTAuthMiddlewareProduct(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Mendapatkan token dari header Authorization
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
            return
        }

        tokenString := parts[1]

        // Memvalidasi token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Pastikan token adalah JWT dengan algoritma yang diharapkan
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            // Gunakan jwtKey yang sama untuk memverifikasi token
            return jwtKey, nil
        })

        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        if !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Jika token valid, lanjutkan ke handler berikutnya
        next.ServeHTTP(w, r)
    })
}

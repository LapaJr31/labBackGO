package main

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Claims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

var jwtKey = []byte("alphabetagamma")
var db *gorm.DB

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Name: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func authenticateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, router *http.Request) {
		tokenString := router.Header.Get("Authorization")

		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, router)
	})
}

func main() {
	db = setupDatabase()

	router := mux.NewRouter()

	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/register", registerHandler).Methods("POST")

	router.HandleFunc("/user", createUserHandler).Methods("POST")
	router.HandleFunc("/user/{id}", getUserHandler).Methods("GET")
	router.HandleFunc("/user/{id}", updateUserHandler).Methods("PUT")
	router.HandleFunc("/user/{id}", deleteUserHandler).Methods("DELETE")

	router.HandleFunc("/category", createCategoryHandler).Methods("POST")
	router.HandleFunc("/category/{id}", getCategoryHandler).Methods("GET")
	router.HandleFunc("/category/{id}", updateCategoryHandler).Methods("PUT")
	router.HandleFunc("/category/{id}", deleteCategoryHandler).Methods("DELETE")

	router.HandleFunc("/expense-record", createExpenseRecordHandler).Methods("POST")
	router.HandleFunc("/expense-record/{id}", getExpenseRecordHandler).Methods("GET")
	router.HandleFunc("/expense-record/{id}", updateExpenseRecordHandler).Methods("PUT")
	router.HandleFunc("/expense-record/{id}", deleteExpenseRecordHandler).Methods("DELETE")

	router.HandleFunc("/currency", createCurrencyHandler).Methods("POST")
	router.HandleFunc("/currency/{id}", getCurrencyHandler).Methods("GET")
	router.HandleFunc("/currency/{id}", updateCurrencyHandler).Methods("PUT")
	router.HandleFunc("/currency/{id}", deleteCurrencyHandler).Methods("DELETE")

	router.HandleFunc("/protected", protectedHandler).Methods("GET")

	log.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

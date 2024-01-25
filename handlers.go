package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func createUserHandler(w http.ResponseWriter, router *http.Request) {
	var user User
	if err := json.NewDecoder(router.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdUser, err := CreateUser(db, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdUser)
}

func getUserHandler(w http.ResponseWriter, router *http.Request) {
	vars := mux.Vars(router)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := GetUser(db, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func updateUserHandler(w http.ResponseWriter, router *http.Request) {
	var user User
	if err := json.NewDecoder(router.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := UpdateUser(db, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteUserHandler(w http.ResponseWriter, router *http.Request) {
	vars := mux.Vars(router)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = DeleteUser(db, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func createCategoryHandler(w http.ResponseWriter, router *http.Request) {
	var category Category
	if err := json.NewDecoder(router.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdCategory, err := CreateCategory(db, category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdCategory)
}

func getCategoryHandler(w http.ResponseWriter, router *http.Request) {
	vars := mux.Vars(router)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	category, err := GetCategory(db, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(category)
}

func getAllCategoriesHandler(w http.ResponseWriter, router *http.Request) {
	categories, err := GetAllCategories(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)
}

func updateCategoryHandler(w http.ResponseWriter, router *http.Request) {
	vars := mux.Vars(router)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var category Category
	if err := json.NewDecoder(router.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	category.ID = uint(id)

	err = UpdateCategory(db, category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteCategoryHandler(w http.ResponseWriter, router *http.Request) {
	vars := mux.Vars(router)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = DeleteCategory(db, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func createExpenseRecordHandler(w http.ResponseWriter, router *http.Request) {
	var record ExpenseRecord
	if err := json.NewDecoder(router.Body).Decode(&record); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdRecord, err := CreateExpenseRecord(db, record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdRecord)
}

func getExpenseRecordHandler(w http.ResponseWriter, router *http.Request) {
	vars := mux.Vars(router)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	record, err := GetExpenseRecord(db, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(record)
}

func updateExpenseRecordHandler(w http.ResponseWriter, router *http.Request) {
	vars := mux.Vars(router)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var record ExpenseRecord
	if err := json.NewDecoder(router.Body).Decode(&record); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	record.ID = uint(id)

	err = UpdateExpenseRecord(db, record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteExpenseRecordHandler(w http.ResponseWriter, router *http.Request) {
	vars := mux.Vars(router)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = DeleteExpenseRecord(db, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func createCurrencyHandler(w http.ResponseWriter, router *http.Request) {
	var currency Currency
	if err := json.NewDecoder(router.Body).Decode(&currency); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdCurrency, err := CreateCurrency(db, currency)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdCurrency)
}

func getCurrencyHandler(w http.ResponseWriter, router *http.Request) {
	vars := mux.Vars(router)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	currency, err := GetCurrency(db, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(currency)
}

func updateCurrencyHandler(w http.ResponseWriter, router *http.Request) {
	vars := mux.Vars(router)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var currency Currency
	if err := json.NewDecoder(router.Body).Decode(&currency); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	currency.ID = uint(id)

	err = UpdateCurrency(db, currency)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteCurrencyHandler(w http.ResponseWriter, router *http.Request) {
	vars := mux.Vars(router)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = DeleteCurrency(db, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func loginHandler(w http.ResponseWriter, router *http.Request) {
	var loginDetails User
	if err := json.NewDecoder(router.Body).Decode(&loginDetails); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user User
	result := db.Where("name = ?", loginDetails.Name).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		}
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDetails.Password))
	if err != nil {
		http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
		return
	}

	tokenString, err := GenerateJWT(strconv.FormatUint(uint64(user.ID), 10))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(tokenString))
}

func protectedHandler(w http.ResponseWriter, router *http.Request) {
	claims, ok := router.Context().Value("claims").(*Claims)
	if !ok {
		http.Error(w, "Error processing JWT claims", http.StatusInternalServerError)
		return
	}

	userID := claims.Name
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Error converting user ID to integer", http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("Access granted for user ID: %d", userIDInt)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func registerHandler(w http.ResponseWriter, router *http.Request) {
	var user User
	if err := json.NewDecoder(router.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error while hashing the password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	createdUser, err := CreateUser(db, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdUser)
}

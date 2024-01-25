package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string `json:"name"`
	Password        string `json:"password"`
	DefaultCurrency string
}

type Category struct {
	gorm.Model
	Name string
}

type ExpenseRecord struct {
	gorm.Model
	UserID     uint
	CategoryID uint
	Timestamp  int64
	Amount     float64
	Currency   string
}

type Currency struct {
	gorm.Model
	Name         string
	ExchangeRate float64 `gorm:"default:1.0"`
}

func setupDatabase() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s ", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	
	db.AutoMigrate(&User{}, &Category{}, &ExpenseRecord{}, &Currency{})

	return db
}

func CreateUser(db *gorm.DB, user User) (User, error) {
	result := db.Create(&user)
	return user, result.Error
}

func GetUser(db *gorm.DB, id uint) (User, error) {
	var user User
	result := db.First(&user, id)
	return user, result.Error
}

func UpdateUser(db *gorm.DB, user User) error {
	return db.Save(&user).Error
}

func DeleteUser(db *gorm.DB, id uint) error {
	return db.Delete(&User{}, id).Error
}

func CreateCategory(db *gorm.DB, category Category) (Category, error) {
	result := db.Create(&category)
	return category, result.Error
}

func GetAllCategories(db *gorm.DB) ([]Category, error) {
	var categories []Category
	result := db.Find(&categories)
	return categories, result.Error
}

func GetCategory(db *gorm.DB, id uint) (Category, error) {
	var category Category
	result := db.First(&category, id)
	return category, result.Error
}

func UpdateCategory(db *gorm.DB, category Category) error {
	return db.Save(&category).Error
}

func DeleteCategory(db *gorm.DB, id uint) error {
	return db.Delete(&Category{}, id).Error
}

func CreateExpenseRecord(db *gorm.DB, record ExpenseRecord) (ExpenseRecord, error) {
	result := db.Create(&record)
	return record, result.Error
}

func GetExpenseRecord(db *gorm.DB, id uint) (ExpenseRecord, error) {
	var record ExpenseRecord
	result := db.First(&record, id)
	return record, result.Error
}

func UpdateExpenseRecord(db *gorm.DB, record ExpenseRecord) error {
	return db.Save(&record).Error
}

func DeleteExpenseRecord(db *gorm.DB, id uint) error {
	return db.Delete(&ExpenseRecord{}, id).Error
}

func CreateCurrency(db *gorm.DB, currency Currency) (Currency, error) {
	result := db.Create(&currency)
	return currency, result.Error
}

func GetCurrency(db *gorm.DB, id uint) (Currency, error) {
	var currency Currency
	result := db.First(&currency, id)
	return currency, result.Error
}

func UpdateCurrency(db *gorm.DB, currency Currency) error {
	return db.Save(&currency).Error
}

func DeleteCurrency(db *gorm.DB, id uint) error {
	return db.Delete(&Currency{}, id).Error
}

package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Credentials stores the database configuration
type Credentials struct {
	Username string
	Password string
	URI      string
	Table    string
}

// DB is a struct that represents a high-level database package
type Database struct {
	credentials Credentials
	Instance    *gorm.DB
}

// Start is a function tha fetches credentials and initiates the connection to the database
func Initiate() *Database {
	fmt.Println("Connecting to Database...")

	// Get the database credentials from the environment
	DB_HOST, _ := os.LookupEnv("DB_HOST")
	DB_USER, _ := os.LookupEnv("DB_USER")
	DB_PASS, _ := os.LookupEnv("DB_PASS")
	DB_NAME, _ := os.LookupEnv("DB_NAME")

	db := &Database{
		credentials: Credentials{
			Username: DB_USER,
			Password: DB_PASS,
			URI:      DB_HOST,
			Table:    DB_NAME,
		},
	}

	db.connect()
	fmt.Println("Database ready for use!")

	return db
}

func (db *Database) connect() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", db.credentials.Username, db.credentials.Password, db.credentials.URI, db.credentials.Table)
	instance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Instance = instance
}

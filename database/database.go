package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	tableName  = "muzz"
	dbUser     = "muzzUser"
	dbPassword = "muzzUserPass"
	uri        = "localhost:3306"
)

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

// Start is a function that starts the database
func Initiate() *Database {
	fmt.Println("Starting Database...")

	db := &Database{
		credentials: Credentials{
			Username: dbUser,
			Password: dbPassword,
			URI:      uri,
			Table:    tableName,
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

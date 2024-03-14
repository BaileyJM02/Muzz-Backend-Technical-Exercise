package context

import (
	"sync"

	"github.com/baileyjm02/muzz-backend-technical-exercise/database"
)

var db *database.Database

// getDatabaseOnce is the database once.
var getDatabaseOnce sync.Once

// GetDB returns the database instance.
func GetDB() *database.Database {
	getDatabaseOnce.Do(func() {
		db = database.Initiate()
	})

	return db
}

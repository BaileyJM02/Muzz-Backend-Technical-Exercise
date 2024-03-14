package context

import (
	"sync"

	"github.com/baileyjm02/muzz-backend-technical-exercise/database"
)

// Context returns the context of the application.
type Context struct {
	// DB is the database package instance.
	DB *database.Database
}

// ctx is the context the application.
var ctx *Context

// ctxOnce ensures the context is only created once.
var ctxOnce sync.Once

// GetContext initiates return the context of the application, ensuring it is only created once.
func GetContext() *Context {
	ctxOnce.Do(func() {
		ctx = &Context{
			DB: GetDB(),
		}
	})

	return ctx
}

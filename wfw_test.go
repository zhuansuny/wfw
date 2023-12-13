package wfw

import (
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	r := New()
	r.Use(Logger())
	api := r.Group("/api")
	api.GET("/test", func(c *Context) {
		c.JSON(http.StatusOK, "Hello world")
	})

	r.Run(":8080")
}

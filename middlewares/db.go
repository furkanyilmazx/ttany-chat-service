package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func DbMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		// before request
		c.Next()
		// after request
	}
}

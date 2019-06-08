package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func LoadRoutes(r *gin.RouterGroup, db *gorm.DB) {
	loadChatRoutes(r, db)
}

package routes;

import (
	"github.com/gin-gonic/gin"
	"ticketingBackend/controllers/auth"
	"ticketingBackend/utils/constants"
	"gorm.io/gorm"
)

func AuthRoutes(router *gin.RouterGroup, db *gorm.DB){

	router.POST(constants.Login, auth.Login(db))

}
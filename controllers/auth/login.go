package auth

import (
	"ticketingBackend/utils"
	"ticketingBackend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)


func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context){

		email:= "useremail@gmail.com"

		tokenInfo:= &models.AuthToken{
			UserEmail: email,
		}

		token, err:= utils.GenerateJWT(tokenInfo)

		if err != nil{
			c.Error(err); // adds/ collects error in the Context which can be processed later
			return;
		}

		c.JSON(http.StatusOK, gin.H{
			"success":true,
			"message":"Login API has been hit",
			"token":token,
		})

	}
}
package functions

import (
	"github.com/gin-gonic/gin"
	"ticketingBackend/models"
)

func GetRequestParameters(c *gin.Context) *models.AuthToken{
	var UserEmail string = ""
	
	if UserEmailValueRaw, okRaw := c.Get("UserEmail"); okRaw {
		if userEmailValue, ok := UserEmailValueRaw.(string); ok{
			UserEmail= userEmailValue;
		} 
	}


	return &models.AuthToken{
		UserEmail: UserEmail,
	}
} 

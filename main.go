package main;

import (
	"github.com/gin-gonic/gin"

	"ticketingBackend/middleware"
	"ticketingBackend/utils/functions"

	"ticketingBackend/database"
	"ticketingBackend/routes"
	"ticketingBackend/utils"

	"github.com/joho/godotenv"
	"log"
	"fmt"
	"net/http"
)

func main(){

	loadEnvError := godotenv.Load()
	if loadEnvError != nil {
		log.Fatal("Error accessing .env file")
	}

	port := utils.GetEnv("PORT")

	dbInstance:=database.InitDb()

	router:=gin.Default()

	router.Use(middleware.RateLimitHandler())
	router.Use(middleware.ErrorHandler())

	router.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"success":true,
			"message":"Server Listening",
		})
	})

	apiGroup:= router.Group("/api")


	authGroup:= apiGroup.Group("/auth")
	routes.AuthRoutes(authGroup, dbInstance)

	privateGroup:= apiGroup.Group("/private")
	privateGroup.Use(middleware.AuthenticationHandler())

	privateGroup.GET("/data", func(c *gin.Context){

		user:= functions.GetRequestParameters(c)

		fmt.Println(user.UserEmail, "reqEmail")


		c.JSON(http.StatusOK, gin.H{
			"success":true,
			"message":"Server Listening",
			"user":user,
		})
	})


	router.Run(":"+ port)
}
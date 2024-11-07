SETUP 
1. go mod init MODULE_NAME

2. go get -u gorm.io/gorm
3. go get -u gorm.io/driver/mysql
4. go get -u github.com/gin-gonic/gin

5. create a .env file in the root directory
6. go get github.com/joho/godotenv // package to lead env variables
7. go get golang.org/x/time/rate // package to handle the rate limit
   Info: exported model name always start with a Capital alphabet
8. go get -u github.com/golang-jwt/jwt/v5

   Info: JSON response {
    c.JSON: This simply sets the response status code and JSON payload and then continues executing the next handler, if there is one. It does not stop further middleware or handlers from running, so it’s useful if you just want to log an error but not prevent the request from completing.

    c.AbortWithStatusJSON: This is typically preferred for error handling in middleware because it sends the response with a JSON payload and status code, then stops any further handlers from executing. This is useful for terminating the request when an error occurs, so that no further processing happens.
   }

KnowMore: {
cannot use "lambaharsh01@gmail.com" (untyped string constant) as *string value in struct literal

This error indicates that you are trying to assign an untyped string constant (like "lambaharsh01@gmail.com") to a field that expects a *string (pointer to a string) rather than a plain string. To fix this, you need to pass a pointer to the string.

Using the & operator directly
email := "lambaharsh01@gmail.com"
yourStruct := YourStruct{
    Email: &email,
}

}


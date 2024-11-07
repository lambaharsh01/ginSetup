package utils

import (
    "time"
    "errors"
    "github.com/golang-jwt/jwt/v5"

	"ticketingBackend/models"
)

func GenerateJWT(tokenInfo *models.AuthToken) (string, error) {
	
	secretKey:= GetEnv("SECRET_KEY")
	var jwtSecret = []byte(secretKey)
	
    claims := jwt.MapClaims{
        "userEmail": tokenInfo.UserEmail,
        "exp": time.Now().Add(time.Hour * 24).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)

}


func ValidateJWT(tokenString string) (*models.AuthToken, error) {

	secretKey:= GetEnv("SECRET_KEY")
	var jwtSecret = []byte(secretKey)
    

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // check if the encription is done with (HS256 || HS384 || HS512) encription standards or not to begain with
            return nil, errors.New("Unexpected signing method")
        }
        return jwtSecret, nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userEmail:= claims["userEmail"].(string)
		authToken:= &models.AuthToken{
			UserEmail : userEmail,
		}
        return authToken, nil
    }

    return nil, errors.New("Invalid token")
}

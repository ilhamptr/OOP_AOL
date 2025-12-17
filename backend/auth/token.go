package auth

import (
	"fmt"
	"net/http"
	"os"
	"product/backend/model"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"crypto/rand"
	"encoding/base64"
	"gorm.io/gorm"
)

type Payload struct{
	ID string
	Email string
	Username string
}


func GenRefreshTkn(db *gorm.DB, user_email string)(string,error){
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate refresh token bytes: %w", err)
	}
	tokenStr := base64.RawURLEncoding.EncodeToString(tokenBytes)
	expiration := time.Now().Add(240 * time.Hour)
	now := time.Now()

	id := uuid.New()
	tkn := models.Session{
		ID: id.String(),
		UserEmail: user_email,
		RefreshToken: tokenStr,
		ExpiresAt: expiration,
		CreatedAt: now,
	}
	
	if err := db.Create(&tkn).Error; err != nil {
		return "", fmt.Errorf("failed to save refresh token: %w", err)
	}
	return tokenStr,nil

}

func Refresh(db *gorm.DB)gin.HandlerFunc{
	return func(c *gin.Context){
		refreshToken := c.GetHeader("refresh_token")
		if refreshToken == ""{
			c.IndentedJSON(http.StatusUnauthorized,gin.H{"error":"Missing refresh token"})
			return
		}

		var tkn models.Session
		var user models.User
		if err := db.Where("refresh_token = ?",refreshToken).First(&tkn).Error; err != nil{
			c.IndentedJSON(http.StatusUnauthorized,gin.H{"error":"Invalid token"})
			return
		}
		if time.Now().Unix() > tkn.ExpiresAt.Unix(){
			c.IndentedJSON(http.StatusUnauthorized,gin.H{"error":"Refresh token expired"})
			return
		}

		if err := db.Where("email = ?",tkn.UserEmail).First(&user).Error; err != nil || tkn.IsRevoked{
			c.IndentedJSON(http.StatusUnauthorized,gin.H{"error":"Invalid token"})
			return
		}

		payloadData := Payload{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		}

		strToken, _ := Token(payloadData)


		c.IndentedJSON(http.StatusOK,gin.H{"access_token":strToken})
	}
}

func Token(data Payload)(string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": data.ID,
		"username": data.Username,
		"email": data.Email,
		// this part set the jwt expiration time
		"exp": time.Now().Add(time.Minute * 10).Unix(),
	})
	tokenSting,err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil{
		return "",err
	}
	return tokenSting,nil

}
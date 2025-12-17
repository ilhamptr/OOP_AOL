package user

import (
	"net/http"
	"product/backend/model"
	"product/backend/service"
	"product/backend/util"
	"product/backend/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"


	"gorm.io/gorm"
)

type SignUp struct{
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginStruct struct{
	Email string `json:"email" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}


func Oauth(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        payload, _ := utils.DecodeSupabaseJWT(c)

        exp, ok := payload["exp"].(float64)
        if !ok || utils.TokenExpiration(exp) {
            c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Access token expired"})
            return
        }

		tokenValid,_ := services.CheckSupabaseToken(c)
		if !tokenValid{
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Access token invalid"})
            return
		}

        email, ok := payload["email"].(string)
        if !ok {
            c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "invalid token email"})
            return
        }
		refreshTkn,_ := auth.GenRefreshTkn(db,email)

        var user models.User
        if err := db.Where("email = ?", email).First(&user).Error; err != nil {

			id := uuid.New()  
			user := models.User{ID: id.String(),Email:email,Username: email}
			db.Create(&user)

			payloadData := auth.Payload{
            ID:       user.ID,
            Email:    user.Email,
            Username: user.Username,
        	}
			
			strToken, err := auth.Token(payloadData)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
				return
			}

            c.IndentedJSON(http.StatusCreated, gin.H{"access_token": strToken,"refresh_token":refreshTkn})
            return
        }

        payloadData := auth.Payload{
            ID:       user.ID,
            Email:    user.Email,
            Username: user.Username,
        }



        strToken, err := auth.Token(payloadData)
        if err != nil {
            c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
            return
        }

        c.IndentedJSON(http.StatusOK, gin.H{"access_token": strToken,"refresh_token":refreshTkn})
    }
}

func VerifyUser(db *gorm.DB)gin.HandlerFunc{
	return func (c *gin.Context)  {
		userId := c.Param("userId")
		var userData models.User
		if err := db.Where("id = ?", userId).First(&userData).Error; err != nil{
			c.IndentedJSON(http.StatusNotFound,gin.H{"error":"account not found"})
			return
		}
		if userData.Active{
			c.IndentedJSON(http.StatusOK,gin.H{"message":"this user is already verified"})
			return
		}
		userData.Active = true
		db.Save(&userData)
		c.IndentedJSON(http.StatusOK,gin.H{"message":"user has been verified successfully"})
	}
}



func Protected(c *gin.Context){
	value,exist := c.Get("claims")
	if !exist {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
	claims := value.(jwt.MapClaims)
	
	userID := claims["id"] 
	c.IndentedJSON(http.StatusAccepted,gin.H{"message": userID})
}



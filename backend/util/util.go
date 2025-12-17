package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
	"product/backend/model"
	"regexp"
	"strings"
	"encoding/json"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"github.com/gin-gonic/gin"
)


func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	specialCharRegex := regexp.MustCompile(`[!@#~$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`)
	return specialCharRegex.MatchString(password)
}


func GenerateOTP() (string,error) {
	max := big.NewInt(1000000)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "",err
	}
	return fmt.Sprintf("%06d", n.Int64()),err
}

func Authenticated(context *gin.Context) (interface{},bool){
	val,exist := context.Get("claims")
	if !exist{
		return val,false
	}
	return val,true
}

func JobAuthorization(userId string, job models.Job) bool{
		return userId == job.AssignBy
}

func DecodeSupabaseJWT(context *gin.Context)(map[string]interface{},error){
	auth := context.GetHeader("Authorization")
	if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
		return nil,fmt.Errorf("missing or invalid Authorization header")
	}

	tokenString := strings.TrimPrefix(auth, "Bearer ")

	parts := strings.Split(tokenString,".")
	if len(parts) < 2{
		return nil,fmt.Errorf("invalid JWT format")

	}
	payloadBytes,_ := base64.RawURLEncoding.DecodeString(parts[1])
	var payload map[string] interface{}

	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
        return nil, fmt.Errorf("invalid JSON payload: %v", err)
    }
	return payload,nil
}

func TokenExpiration(exp float64) bool{
	return time.Now().Unix() > int64(exp)
}


func VerifyToken(tokenString string)(jwt.MapClaims,error){
	token,err := jwt.Parse(tokenString,func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")),nil
	})
	if err != nil{
		return nil,err
	}
	if !token.Valid{
		return nil,fmt.Errorf("invalid token")
	}
	claims,ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil,fmt.Errorf("invalid token claim")
	}
	return claims,nil
}

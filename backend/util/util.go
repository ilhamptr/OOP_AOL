package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"
	"github.com/gin-gonic/gin"
	"product/backend/model"
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
package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
    "gorm.io/gorm"
	"github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "product/backend/model"
    "product/backend/util"
)

func GetResumeFile(db *gorm.DB)gin.HandlerFunc{
    return func(c *gin.Context) {
        resumeName := c.Param("resumeName")
        val,exist := utils.Authenticated(c)

        if !exist{
			c.IndentedJSON(http.StatusBadRequest,gin.H{"error":"authentication failed, please try again"})
			return
        }

		claims := val.(jwt.MapClaims)
		userId := claims["id"].(string)

        var applicant models.Application
        if err := db.First(&applicant,"resume_file = ?",resumeName).Error;err != nil{
            c.IndentedJSON(http.StatusNotFound,gin.H{"error":"record not found"})
			return 
        }

        var job models.Job
		if err := db.First(&job,"id = ?",applicant.JobID).Error; err != nil{
			c.IndentedJSON(http.StatusNotFound,gin.H{"error":"record not found"})
			return 
		}

        if userId != job.AssignBy{
            c.IndentedJSON(http.StatusUnauthorized,gin.H{"error":"you are not authorize to access this file"})
			return 
        }
        url, err := generateSignedURL(resumeName)
        if err != nil{
            c.IndentedJSON(http.StatusInternalServerError,gin.H{"error":"an unexpected error has occurred"})
			return 
        }
        c.IndentedJSON(http.StatusOK, gin.H{
            "url": url,
        })
    }
}

func generateSignedURL(fileName string) (string, error) {
	projectID := os.Getenv("PROJECT_ID")
	bucket := os.Getenv("BUCKET")
	token := os.Getenv("SUPABASE_TOKEN")
    signURL := fmt.Sprintf(
        "https://%s.supabase.co/storage/v1/object/sign/%s/%s",
        projectID,
        bucket,
        fileName,
    )

    reqBody := `{"expiresIn": 3600}` 

    req, err := http.NewRequest("POST", signURL, bytes.NewBuffer([]byte(reqBody)))
    if err != nil {
        return "", err
    }
    req.Header.Set("Authorization", "Bearer "+token)
    req.Header.Set("Content-Type", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)

    var result map[string]interface{}
    json.Unmarshal(body, &result)

    if signedPath, ok := result["signedURL"].(string); ok {
        return fmt.Sprintf("https://%s.supabase.co/storage/v1/%s", projectID, signedPath), nil
    }

    return "", fmt.Errorf("failed to generate signed URL: %s", body)
}

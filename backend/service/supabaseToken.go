package services

import(
	"net/http"
	"time"
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"strings"
)


func CheckSupabaseToken(context *gin.Context) (bool, error) {
	auth := context.GetHeader("Authorization")
	if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
		return false,fmt.Errorf("missing or invalid Authorization header")
	}

	tokenString := strings.TrimPrefix(auth, "Bearer ")
	supabaseURL := os.Getenv("SUPABASE_URL")
	apiKey:= os.Getenv("SUPABASE_TOKEN_VALIDATION_API_KEY")

	url := fmt.Sprintf("%s/auth/v1/user",supabaseURL)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+tokenString)
	req.Header.Set("apikey", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK, nil
}
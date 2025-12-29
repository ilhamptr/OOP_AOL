package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
	"gorm.io/gorm"
	"product/backend/model"
	"sync"


)

type ScoringInput struct{
	JobDesc string `json:"jobDescription" binding:"required"`
	TopNumber int `json:"topNumber" binding:"required"`
}

type EvaluationInput struct{
	JobDesc string `json:"jobDescription" binding:"required"`
	ResumeName string `json:"resumeName" binding:"required"`
}

type ApplicantDataRes struct{
	Id string `json:"id"`
	ApplicantName string `json:"applicant_name"`
	ResumeName string `json:"resume_file"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	AppliedAt  time.Time `json:"created_at"`
}

func SubmitResume(fileBytes []byte,jobID,applicantName,filename string) ([]byte,error) {
	projectID := os.Getenv("AI_API_ENDPOINT")
	endpoint := fmt.Sprintf("%v/resume-processing/%v",projectID,jobID)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fileField,err := writer.CreateFormFile("file",filename)
	if err != nil{
		return nil,err
	}
	if _, err := fileField.Write(fileBytes); err != nil {
        return nil,err
    }
	_ = writer.WriteField("applicantName",applicantName)
	_ = writer.WriteField("resumeFile",filename)
	writer.Close()
	req,_ := http.NewRequest("POST",endpoint,body)
	req.Header.Set("Content-Type",writer.FormDataContentType())
	client := &http.Client{}
	resp,err := client.Do(req)
	if err != nil{
		return nil,err
	}
	defer resp.Body.Close()
	respBody,_ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusCreated{
		return nil, fmt.Errorf("upload failed: %s", resp.Status)
	}
	return respBody,nil

}

func GetApplicantData(db *gorm.DB,resumeName string) (map[string]interface{},error){
	var applicant models.Application
	if err := db.Where("resume_file = ?",resumeName).First(&applicant).Error; err != nil {
		return nil,err
	}
	resp := ApplicantDataRes{
		Id: applicant.ID,
		ApplicantName: applicant.ApplicantName,
		ResumeName: applicant.ResumeFile,
		Email: applicant.Email,
		Phone: applicant.PhoneNumber,
		AppliedAt: applicant.CreatedAt,
	}

	result := map[string]interface{}{
		"id":             resp.Id,
		"applicant_name": resp.ApplicantName,
		"resume_file":    resp.ResumeName,
		"email":          resp.Email,
		"phone":          resp.Phone,
		"created_at":     resp.AppliedAt,
	}

	return result, nil

}


func GetScoring(db *gorm.DB,token,jobDesc,jobId string, topNumber int) (map[string]interface{},error){
	projectID := os.Getenv("AI_API_ENDPOINT")
	endpoint := fmt.Sprintf("%v/get-applicants/%v",projectID,jobId)

	data := ScoringInput{
		JobDesc: jobDesc,
		TopNumber: topNumber,
	}

	jsonData,_ := json.Marshal(data)
	req,_ := http.NewRequest("POST",endpoint,bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	client := &http.Client{}
	resp,err:= client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes,_ := io.ReadAll(resp.Body)
	var result map[string]interface{}

	if resp.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("upload failed: %s", resp.Status)
	}

	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return nil, err
	}

	dataArr, ok := result["data"].([]interface{})
	count := len(dataArr)

	if !ok || count == 0 {
		return nil, fmt.Errorf("data field missing or empty")
	}
	results := make([]map[string]interface{}, count)

	var wg sync.WaitGroup
	wg.Add(count)

	for i, raw := range dataArr {
		go func(index int, item interface{}) {
			defer wg.Done()

			obj, ok := item.(map[string]interface{})
			if !ok {
				return
			}

			fileName, ok := obj["resume_file"].(string)
			if !ok {
				return
			}

			applicantData, err := GetApplicantData(db, fileName)
			if err != nil {
				return
			}

			results[index] = applicantData  
		}(i, raw)
	}

	wg.Wait()

	// remove nil entries if any goroutine failed
	final := make([]map[string]interface{}, 0, count)
	for _, r := range results {
		if r != nil {
			final = append(final, r)
		}
	}

	response := map[string]interface{}{
		"data": final,
	}

	return response, nil
	}

func ScoringDetails(token,jobDesc,resumeName string) (map[string]interface{},error){
	projectID := os.Getenv("AI_API_ENDPOINT")
	endpoint := fmt.Sprintf("%v/matching-process/",projectID)

	data := EvaluationInput{
		JobDesc: jobDesc,
		ResumeName: resumeName,
	}

	jsonData,_ := json.Marshal(data)
	req,_ := http.NewRequest("POST",endpoint,bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	client := &http.Client{}
	resp,err:= client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes,_ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	if resp.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("upload failed: %s", resp.Status)
	}

	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return nil, err
	}

	return result,nil
}
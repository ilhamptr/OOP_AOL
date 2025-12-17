package main

import (
	"product/backend/auth"
	"product/backend/job"
	"product/backend/middleware"
	"product/backend/model"
	"product/backend/service"
	"product/backend/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	router.Use(cors.New(cors.Config{
        AllowAllOrigins: true,
        AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:    []string{"Origin", "Content-Type", "Authorization", "refresh_token"},
        ExposeHeaders:   []string{"Content-Length"},
    }))


	db,_ := models.DbSession()

	// protected Endpoints
	protectedEnpoints := router.Group("")
	protectedEnpoints.Use(middleware.Authentication)
	protectedEnpoints.POST("/protected",user.Protected)
	protectedEnpoints.POST("/add-job",job.AddJob(db))
	protectedEnpoints.GET("/jobs",job.GetJobs(db))
	protectedEnpoints.DELETE("/delete-job/:jobId",job.DeleteJob(db))
	protectedEnpoints.GET("/view-applicants/:jobId",job.SeeApplicants(db))
	protectedEnpoints.POST("/view-top-applicants/:jobId",job.ScoringApplicants(db))
	protectedEnpoints.POST("/view-applicant-evaluation/:jobId/:resumeName",job.ApplicantDetails(db))
	protectedEnpoints.GET("/download-resume/:resumeName",services.GetResumeFile(db))


	router.POST("/apply/:jobId",job.Apply(db))
	router.GET("/job-details/:jobId",job.GetJobDetails(db))
	router.POST("/auth",user.Oauth(db))
	router.POST("/refresh",auth.Refresh(db))
	router.GET("/verify-account/:userId",user.VerifyUser(db))
	router.Run("localhost:8080")
}

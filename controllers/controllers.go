package controllers

import (
	"job-app/main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var jobs = []models.Job{
	{ID: "1", Title: "google", Description: "testssfds ", Status: "applied"},
}

func GetJobs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": jobs})
}

func AddBook(c *gin.Context) {
	var newJob models.Job
	if err := c.BindJSON(&newJob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobs = append(jobs, newJob)

	c.JSON(http.StatusCreated, gin.H{"data": newJob})
}

func EditJob(c *gin.Context) {
	var editJob models.Job
	if err := c.BindJSON(&editJob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jobId := c.Param("id")

	for i, job := range jobs {
		if job.ID == jobId {
			if editJob.Title != "" {
				jobs[i].Title = editJob.Title
			}
			if editJob.Description != "" {
				jobs[i].Description = editJob.Description
			}
			if editJob.Status != "" {
				jobs[i].Status = editJob.Status
			}
			c.JSON(http.StatusOK, jobs[i])
			return

		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "job not found"})
}

func DeleteJob(c *gin.Context) {

	jobId := c.Param("id")

	for i, job := range jobs {
		if job.ID == jobId {
			jobs = append(jobs[:i], jobs[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "job not found"})
}

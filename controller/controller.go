package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardosoalves/api-go-gin/controller/request"
	"github.com/renatocardosoalves/api-go-gin/controller/response"
	"github.com/renatocardosoalves/api-go-gin/database"
	"github.com/renatocardosoalves/api-go-gin/model"
)

func GetAll(c *gin.Context) {
	var students []model.Student
	database.DB.Find(&students)
	var res []response.StudentResponse
	for _, student := range students {
		res = append(res, response.StudentResponse{
			ID:        student.ID,
			FirstName: student.FirstName,
			LastName:  student.LastName,
			Email:     student.Email,
		})
	}
	c.JSON(http.StatusOK, res)
}

func GetByID(c *gin.Context) {
	var student model.Student
	database.DB.First(&student, c.Params.ByName("id"))

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}

	res := response.StudentResponse{
		ID:        student.ID,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Email:     student.Email,
	}
	c.JSON(http.StatusOK, res)
}

func Create(c *gin.Context) {
	var req request.StudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := model.Student{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	database.DB.Create(&student)

	res := response.StudentResponse{
		ID:        student.ID,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Email:     student.Email,
	}
	c.JSON(http.StatusCreated, res)
}

func Delete(c *gin.Context) {
	var student model.Student
	if result := database.DB.Delete(&student, c.Params.ByName("id")); result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func Update(c *gin.Context) {
	var req request.StudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var student model.Student
	if result := database.DB.First(&student, c.Params.ByName("id")); result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}
	student.FirstName = req.FirstName
	student.LastName = req.LastName
	student.Email = req.Email
	database.DB.Model(&student).Updates(student)
	res := response.StudentResponse{
		ID:        student.ID,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Email:     student.Email,
	}
	c.JSON(http.StatusOK, res)
}

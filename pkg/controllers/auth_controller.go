package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapildev333/go_safety_plan/pkg/models"
)

func RegisterUser(c *gin.Context){
var input models.RegisterUserModel

	if err := c.ShouldBindJSON(&input); 
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	_,err := u.SaveUser()

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"registration success"})
}
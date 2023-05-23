package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kapildev333/go_safety_plan/pkg/controllers"
	"github.com/kapildev333/go_safety_plan/pkg/models"
)

func main(){
	models.ConnectDataBase()
	r := gin.Default()
	
	public := r.Group("/api")

	public.POST("/register", controllers.RegisterUser)

	r.Run(":8080")
}
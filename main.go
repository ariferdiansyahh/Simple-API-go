package main

import (
	cakecontroller "github.com/ariferdiansyah/ralali_backend/controllers/cakeController"
	"github.com/ariferdiansyah/ralali_backend/models"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	models.ConnectDatabase()
	r.GET("/api/cakes", cakecontroller.Index)
	r.GET("/api/cakes/:id", cakecontroller.Show)
	r.POST("/api/cakes", cakecontroller.Create)
	r.PUT("/api/cakes/:id", cakecontroller.Update)
	r.DELETE("/api/cakes/:id", cakecontroller.Delete)

	r.Run()
}

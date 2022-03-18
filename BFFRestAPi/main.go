package main

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"rest_grpc_test/BFFRestAPI/client"
)

var (
	timeout       = time.Second
	user_client   client.UsersClient
)

func UserRegister(router *gin.RouterGroup) {
	router.GET("/", GetUsers)
}

func GetUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	data, err := user_client.GreetUser(&ctx)
	response(c, data, err)
}

func response(c *gin.Context, data interface{}, err error) {
	statusCode := http.StatusOK
	var errorMessage string
	if err != nil {
		log.Println("Server Error Occured:", err)
		errorMessage = strings.Title(err.Error())
		statusCode = http.StatusInternalServerError
	}
	c.JSON(statusCode, gin.H{"data": data, "error": errorMessage})
}

func main() {
	log.Println("Bff Service")

	r := gin.Default()
	r.Use(cors.Default())

	api := r.Group("/api")
	UserRegister(api.Group("/users"))

	r.Run()
}

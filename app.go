package main

import (
	"os"

	"./auth"
	"./global"
	"./mail"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	err := global.InitDatabase()
	if err != nil {
		os.Exit(1)
	}

	r.POST("/login", auth.Login)
	r.POST("/register", auth.Register)

	useCredential := r.Group("/")
	useCredential.Use(
		auth.CheckToken(),
	)
	{
		useCredential.POST("/send_email", mail.SendMailToNode)
	}

	r.Run()
}

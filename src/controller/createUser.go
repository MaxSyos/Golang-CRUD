package controller

import (
	"fmt"
	"log"

	"github.com/MaxSyos/Golang-CRUD/src/config/validation"
	"github.com/MaxSyos/Golang-CRUD/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Erro na criação de usuário, error=%s", err.Error())
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}

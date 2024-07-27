package controller

import (
	"fmt"

	"github.com/MaxSyos/Golang-CRUD/src/config/rest_err"
	"github.com/MaxSyos/Golang-CRUD/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("Dados incorretos na criação do usuário, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}

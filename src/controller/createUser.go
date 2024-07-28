package controller

import (
	"fmt"

	"github.com/MaxSyos/Golang-CRUD/src/config/logger"
	"github.com/MaxSyos/Golang-CRUD/src/config/validation"
	"github.com/MaxSyos/Golang-CRUD/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro na criação de usuário, error=%s", err,
			zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}

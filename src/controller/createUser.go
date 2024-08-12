package controller

import (
	"net/http"

	"github.com/MaxSyos/Golang-CRUD/src/config/logger"
	"github.com/MaxSyos/Golang-CRUD/src/config/validation"
	"github.com/MaxSyos/Golang-CRUD/src/controller/model/request"
	"github.com/MaxSyos/Golang-CRUD/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
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

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Usuário Criado com Sucesso", zap.String("journey", "createUser"))
	c.String(http.StatusOK, "")

}

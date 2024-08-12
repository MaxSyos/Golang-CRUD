package model

import (
	"fmt"

	"github.com/MaxSyos/Golang-CRUD/src/config/logger"
	"github.com/MaxSyos/Golang-CRUD/src/config/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Iniciando a criação do Usuário", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}

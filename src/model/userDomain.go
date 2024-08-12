package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/MaxSyos/Golang-CRUD/src/config/rest_err"
)

type userDomain struct {
	email    string
	password string
	name     string
	age      int8
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email, password, name, age,
	}
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*userDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}

/* func (ud *userDomain) GetID() string {
	return ud.id
}

func (ud *userDomain) SetID(id string) {
	ud.id = id
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetAge() int8 {
	return ud.age
} */

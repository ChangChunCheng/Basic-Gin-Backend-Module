package api

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"

	"basic-gin-backend-module/handler"
	"basic-gin-backend-module/model"
)

type User struct {
	UserID   string    `json:"UserID"`
	Account  string    `json:"Account"`
	Name     string    `json:"Name"`
	Disable  bool      `json:"Disable"`
	CreateAt time.Time `json:"CreateAt"`
	UpdateAt time.Time `json:"UpdateAt"`
}

func (user *User) FromModel(userModel *model.Users) {
	user.UserID = userModel.UserID.String()
	user.Account = userModel.Account
	user.Name = userModel.Name
	user.Disable = userModel.Disable
	user.CreateAt = userModel.CreateAt
	user.UpdateAt = userModel.UpdateAt
}

type CreateUser struct {
	Account  string `json:"Account"`
	Password string `json:"Password"`
	Name     string `json:"Name"`
}

func (createUser *CreateUser) ToModel(userModel *model.Users, err error) {
	secret, err := handler.GetSecret(createUser.Password)
	if err != nil {
		panic(err)
	}
	userModel.Account = createUser.Account
	userModel.Password = []byte(secret)
	userModel.Name = createUser.Name
}

type UpdateUser struct {
	UserID   string `json:"UserID"`
	Account  string `json:"Account"`
	Password string `json:"Password"`
	Name     string `json:"Name"`
}

func (updateUser *UpdateUser) ToModel(userModel *model.Users, err error) {
	secret, err := handler.GetSecret(updateUser.Password)
	if err != nil {
		panic(err)
	}
	userModel.UserID = uuid.FromStringOrNil(updateUser.UserID)
	userModel.Account = updateUser.Account
	userModel.Password = []byte(secret)
	userModel.Name = updateUser.Name
}

type DisableUser struct {
	UserID string `json:"UserID"`
}

func (disableUser *DisableUser) ToModel(userModel *model.Users) {
	userModel.UserID = uuid.FromStringOrNil(disableUser.UserID)
}

type EnableUser struct {
	UserID string `json:"UserID"`
}

func (enableUser *EnableUser) ToModel(userModel *model.Users) {
	userModel.UserID = uuid.FromStringOrNil(enableUser.UserID)
}

type DeleteUser struct {
	UserID string `json:"UserID"`
}

func (deleteUser *DeleteUser) ToModel(userModel *model.Users) {
	userModel.UserID = uuid.FromStringOrNil(deleteUser.UserID)
}

// Login - Login data to BindJSON()
type Login struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

// Claims - Token data to build JWT
type Claims struct {
	UserID   int32  `json:"userid"`
	Account  string `json:"account"`
	Name     string `json:"name"`
	Ordered  bool   `json:"ordered"`
	Identity string `json:"identity"`
	jwt.StandardClaims
}

package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"basic-gin-backend-module/dao"
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

func (createUser *CreateUser) ToModel(userModel *model.Users) error {
	secret, err := handler.GetSecret(createUser.Password)
	if err != nil {
		return err
	}
	userModel.Account = createUser.Account
	userModel.Password = []byte(secret)
	userModel.Name = createUser.Name
	return nil
}

type UpdateUser struct {
	UserID   string `json:"UserID"`
	Password string `json:"Password"`
	Name     string `json:"Name"`
}

func (updateUser *UpdateUser) ToModel(userModel *model.Users) error {
	secret, err := handler.GetSecret(updateUser.Password)
	if err != nil {
		return err
	}
	userModel.UserID = uuid.FromStringOrNil(updateUser.UserID)
	userModel.Password = []byte(secret)
	userModel.Name = updateUser.Name
	return nil
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

func getMe(c *gin.Context) {
	claims := c.MustGet("claims").(Claims)

	userRes := &User{
		UserID:  claims.UserID,
		Account: claims.Account,
		Name:    claims.Name,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userRes,
	})
}

func updateUser(c *gin.Context) {
	var updateUser *UpdateUser
	var err error
	if err = c.BindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": RequestBodyError,
		})
		c.Abort()
		return
	}

	userModel := &model.Users{}
	err = updateUser.ToModel(userModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": ProcessError,
		})
		c.Abort()
		return
	}

	user, err := dao.AddUsers(userModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": DBInsertError,
		})
		c.Abort()
		return
	}

	userRes := &User{}
	userRes.FromModel(user)

	c.JSON(http.StatusOK, gin.H{
		"data": userRes,
	})
}

func configUserRouter(group string, e *gin.Engine) {
	g := e.Group(group)
	g.Use(auth)
	g.PUT("/update", updateUser)
	g.GET("/me", getMe)
}

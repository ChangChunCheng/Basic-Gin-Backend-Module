package api

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"basic-gin-backend-module/dao"
	"basic-gin-backend-module/model"
)

// Login - Login data to BindJSON()
type Login struct {
	Account  string `json:"Account"`
	Password string `json:"Password"`
}

// Claims - Token data to build JWT
type Claims struct {
	UserID  string
	Account string
	Name    string
	jwt.StandardClaims
}

func signup(c *gin.Context) {
	var createUser *CreateUser
	var err error
	if err = c.BindJSON(&createUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": RequestBodyError,
		})
		c.Abort()
		return
	}

	userModel := &model.Users{}
	err = createUser.ToModel(userModel)
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

	c.JSON(http.StatusCreated, gin.H{
		"data": userRes,
	})
}

func login(c *gin.Context) {
	var login *Login
	var err error
	if err = c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": RequestBodyError,
		})
		c.Abort()
		return
	}

	user, err := dao.GetUsersByAccount(login.Account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": AccountNotExist,
		})
		c.Abort()
		return
	}

	claims := &Claims{
		UserID:  user.UserID.String(),
		Account: user.Account,
		Name:    user.Name,
	}

	token, err := GetJWT(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": ProcessError,
		})
		c.Abort()
		return
	}

	userRes := &User{}
	userRes.FromModel(user)

	c.JSON(http.StatusOK, gin.H{
		"data":  userRes,
		"token": token,
	})
}

func auth(c *gin.Context) {
	auth := c.GetHeader(viper.GetString("auth.header"))
	claims, message := ParseJWT(auth)
	// fmt.Println(message);
	if message == "" {
		c.Set("claims", claims)
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": Unauthorize,
		})
		c.Abort()
	}
}

func configAuthRouter(group string, e *gin.Engine) {
	g := e.Group(group)
	g.POST("/signup", signup)
	g.POST("/login", login)
}

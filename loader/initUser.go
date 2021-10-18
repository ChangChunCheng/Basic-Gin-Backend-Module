package loader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"reflect"

	"basic-gin-backend-module/api"
	"basic-gin-backend-module/dao"
	"basic-gin-backend-module/model"
)

func initUser() {
	var err error

	data, err := ioutil.ReadFile(path.Join("config", "initUser.json"))
	if err != nil {
		panic("InitUser reading error")
	}

	var createUser *api.CreateUser
	err = json.Unmarshal(data, &createUser)
	if err != nil {
		panic("InitUser json unmarshal error")
	}

	createUserModel := &model.Users{}
	createUser.ToModel(createUserModel, err)
	if err != nil {
		panic("InitUser user secret error")
	}

	userModel, err := dao.GetUsersByAccount(createUserModel.Account)
	if err != nil {
		fmt.Println("InitUser with the default password")
	} else {
		if reflect.DeepEqual(createUser.Password, userModel.Password) {
			fmt.Println("InitUser with the default password")
		} else {
			fmt.Println("InitUser without the default password")
		}
		return
	}

	_, rows, err := dao.AddUsers(createUserModel)
	if (err != nil) || (rows != 1) {
		panic("Init user error")
	}
}

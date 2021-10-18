package dao

import (
	"time"

	null "github.com/guregu/null"
	uuid "github.com/satori/go.uuid"

	"basic-gin-backend-module/model"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

func GetUsersByAccount(argAccount string) (record *model.Users, err error) {
	record = &model.Users{}
	if err = DB.Where("Account = ? ", argAccount).First(record).Error; err != nil {
		err = ErrNotFound
		return nil, err
	}
	return record, nil
}

func AddUsers(record *model.Users) (result *model.Users, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

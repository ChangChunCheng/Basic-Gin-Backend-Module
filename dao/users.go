package dao

import (
	// "sync"
	"time"

	null "github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"

	"basic-gin-backend-module/model"
)

var (
	// wg sync.WaitGroup
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

func GetUsers(argUserID uuid.UUID) (record *model.Users, err error) {
	record = &model.Users{}
	if err = DB.First(record, argUserID).Error; err != nil {
		err = ErrNotFound
		return nil, err
	}
	return record, nil
}

func GetUsersByAccount(argAccount string) (record *model.Users, err error) {
	record = &model.Users{}
	if err = DB.Where("Account = ? ", argAccount).First(record).Error; err != nil {
		err = ErrNotFound
		return nil, err
	}
	return record, nil
}

func AddUsers(record *model.Users) (result *model.Users, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(record).Error; err != nil {
			return ErrInsertFailed
		}
		return nil
	})

	return record, err
}

func UpdateUsers(record *model.Users) (result *model.Users, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		result = &model.Users{}
		result, err := GetUsers(record.UserID)
		if err != nil {
			return ErrNotFound
		}
		if err = Copy(result, record); err != nil {
			return ErrUpdateFailed
		}
		result.Disable = record.Disable

		if err = tx.Save(result).Error; err != nil {
			return ErrUpdateFailed
		}
		return nil
	})

	return result, err
}

func DeleteUser(argUserID uuid.UUID) (err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		result, err := GetUsers(argUserID)
		if err != nil {
			return ErrNotFound
		}
		if err = tx.Delete(result.UserID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

package model

import (
	"database/sql"
	"time"

	null "github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

type Users struct {
	UserID   uuid.UUID      `gorm:"primary_key;column:userid;type:UUID;default:GEN_RANDOM_UUID();"`
	Account  string         `gorm:"column:account;type:VARCHAR;size:50;"`
	Password []byte         `gorm:"column:password;type:BYTEA;"`
	Name     string         `gorm:"column:name;type:VARCHAR;size:50"`
	Disable  bool           `gorm:"column:disable;type:BOOL;default:false;"`
	CreateAt time.Time      `gorm:"column:createat;type:TIMESTAMP;default:time.Now();"`
	UpdateAt time.Time      `gorm:"column:updateat;type:TIMESTAMP;default:time.Now();"`
	DeleteAt gorm.DeletedAt `gorm:"column:deleteat;type:TIMESTAMP;"`
}

var usersTableInfo = &TableInfo{
	Name: "users",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "userid",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "UUID",
			DatabaseTypePretty: "UUID",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "UUID",
			ColumnLength:       -1,
			GoFieldName:        "Userid",
			GoFieldType:        "uuid.UUID",
		},

		&ColumnInfo{
			Index:              1,
			Name:               "account",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       50,
			GoFieldName:        "Account",
			GoFieldType:        "string",
		},

		&ColumnInfo{
			Index:              2,
			Name:               "password",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BYTEA",
			DatabaseTypePretty: "BYTEA",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BYTEA",
			ColumnLength:       -1,
			GoFieldName:        "Password",
			GoFieldType:        "[]byte",
		},

		&ColumnInfo{
			Index:              3,
			Name:               "disable",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "disable",
			GoFieldType:        "bool",
		},

		&ColumnInfo{
			Index:              5,
			Name:               "create_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "CreateAt",
			GoFieldType:        "time.Time",
		},

		&ColumnInfo{
			Index:              6,
			Name:               "update_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "UpdateAt",
			GoFieldType:        "time.Time",
		},

		&ColumnInfo{
			Index:              7,
			Name:               "delete_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "DeleteAt",
			GoFieldType:        "time.Time",
		},
	},
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *Users) BeforeSave(*gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *Users) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *Users) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *Users) TableInfo() *TableInfo {
	return usersTableInfo
}

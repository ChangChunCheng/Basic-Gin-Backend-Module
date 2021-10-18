// Package loader
package loader

import (
	// "os"
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"basic-gin-backend-module/dao"
	"basic-gin-backend-module/model"
)

// GORM executing configuration
var (
	dbUser    string = mustGetenv("DB_USER")
	dbPwd     string = mustGetenv("DB_PASS")
	dbTCPHost        = mustGetenv("DB_HOST")
	// instanceConnectionName string = mustGetenv("INSTANCE_CONNECTION_NAME")
	dbName string = mustGetenv("DB_NAME")

	dbPort string = os.Getenv("DB_PORT")

	MaxLifetime  int = 10
	MaxOpenConns int = 10
	MaxIdleConns int = 10
)

// loadPSQL - Loading PSQL by gorm
func loadPSQL() *gorm.DB {

	dbURI := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s", dbTCPHost, dbUser, dbPwd, dbPort, dbName)

	// 2. Open gorm connection
	var conn *gorm.DB
	var err error
	if viper.GetString("app.mode") == "develop" || viper.GetString("app.mode") == "test" {
		conn, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{
			Logger: logger.New(
				dbLogger,
				logger.Config{
					SlowThreshold: time.Millisecond,
					// Silent, Error, Warn, Info
					LogLevel: logger.Info,
				},
			),
			PrepareStmt: true,
		})
	} else {
		conn, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{FullSaveAssociations: true})
	}
	if err != nil {
		panic(fmt.Errorf("psql connection error"))
	}

	// !Important
	// 3. Assign connection to DB CRUD (dao)
	dao.DB = conn

	// 4. Try to connect DB
	db, err := conn.DB()
	if err != nil {
		panic(fmt.Errorf("PSQL loadding error"))
	}

	// 5. Setting db conf
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(MaxLifetime))

	// 6. Migrating the gorm.Model to DB tables
	conn.AutoMigrate(

		// Users
		&model.Users{},
	)
	fmt.Println("PSQL AutoMigrate finished.")

	return conn
}

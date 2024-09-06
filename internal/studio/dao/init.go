package dao

import (
	"TTMS/configs/consts"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Studio struct {
	Id        int64
	Name      string
	RowsCount int64
	ColsCount int64
}

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(consts.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,                                //禁用默认事务操作
			Logger:                 logger.Default.LogMode(logger.Info), //打印sql语句
		},
	)
	if err != nil {
		panic(err)
	}
	sqlDB, _ := DB.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
}

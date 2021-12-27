package mysql

import (
	"server-grpc/application/domain/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlConnection(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database " + err.Error())
	}

	db.AutoMigrate(&entity.MovieSearchLog{})
	return db
}

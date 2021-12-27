package mysql

import (
	"server-grpc/application/domain/entity"
	"server-grpc/application/domain/repository"
	"time"

	"gorm.io/gorm"
)

type movieMysqlRepositoryImpl struct {
	db *gorm.DB
}

func NewMovieMysqlRepository(db *gorm.DB) repository.MovieMysqlRepository {
	return &movieMysqlRepositoryImpl{
		db: db,
	}
}

func (repo *movieMysqlRepositoryImpl) CreateLogSearch(keywordSearch string) error {
	return repo.db.Create(&entity.MovieSearchLog{
		KeywordSearch: keywordSearch,
		CreatedAt:     time.Now(),
	}).Error
}

package repository

import (
	"server-grpc/application/domain/entity"
)

type MovieOmdbRepository interface {
	GetListMovie(*entity.MovieFilter) (*entity.ListMovie, error)
	GetMovieDetail(ImdbId string) (*entity.MovieDetail, error)
}

type MovieMysqlRepository interface {
	CreateLogSearch(keywordSearch string) error
}

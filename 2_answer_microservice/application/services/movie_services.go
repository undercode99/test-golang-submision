package services

import (
	"server-grpc/application/domain/entity"
	"server-grpc/application/domain/repository"
	"server-grpc/config"
	"server-grpc/persistence/mysql"
	"server-grpc/persistence/omdb"
)

func NewMovieService(repoOmdb repository.MovieOmdbRepository, repoMysql repository.MovieMysqlRepository) *MovieServices {
	return &MovieServices{
		movieOmdbRepository:  repoOmdb,
		movieMysqlRepository: repoMysql,
	}
}

func SetupMovieServices() *MovieServices {
	db := mysql.NewMysqlConnection(config.MysqlDsn())
	movieMysqlRepo := mysql.NewMovieMysqlRepository(db)
	movieOmdbRepo := omdb.NewMovieOmdbRepository(config.OmdbApiKey(), config.OmdbApiUrl())
	movieService := NewMovieService(movieOmdbRepo, movieMysqlRepo)
	return movieService
}

type MovieServices struct {
	movieOmdbRepository  repository.MovieOmdbRepository
	movieMysqlRepository repository.MovieMysqlRepository
}

func (service *MovieServices) GetListMovie(filter *entity.MovieFilter) (*entity.ListMovie, error) {
	if filter.Search != "" {
		go service.movieMysqlRepository.CreateLogSearch(filter.Search)
	}
	return service.movieOmdbRepository.GetListMovie(filter)
}

func (service *MovieServices) GetMovieDetail(imdbID string) (*entity.MovieDetail, error) {
	return service.movieOmdbRepository.GetMovieDetail(imdbID)
}

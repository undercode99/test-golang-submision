package omdb

import (
	"encoding/json"
	"net/http"
	"server-grpc/application/domain/entity"
	"server-grpc/application/domain/repository"
	"strconv"
)

type movieOmdbRepositoryImpl struct {
	apiKey string
	apiUrl string
}

func NewMovieOmdbRepository(apiKey string, apiUrl string) repository.MovieOmdbRepository {
	return &movieOmdbRepositoryImpl{
		apiKey: apiKey,
		apiUrl: apiUrl,
	}
}

func (repo *movieOmdbRepositoryImpl) requestApiUrl() string {
	return repo.apiUrl + "?apikey=" + repo.apiKey
}

func (repo *movieOmdbRepositoryImpl) GetListMovie(filter *entity.MovieFilter) (*entity.ListMovie, error) {

	requestApiUrl := repo.requestApiUrl() + "&s=" + filter.Search

	if filter.Page > 0 {
		requestApiUrl += "&page=" + strconv.Itoa(filter.Page)
	}

	listMovie := &entity.ListMovie{}

	response, err := http.Get(requestApiUrl)
	if err != nil {
		listMovie.Response = "False"
		listMovie.ErrorMessage = "Error when request api"
		return listMovie, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(listMovie)
	if err != nil {
		listMovie.Response = "False"
		listMovie.ErrorMessage = "Error when decode response"
		return listMovie, err
	}

	return listMovie, nil
}

func (repo *movieOmdbRepositoryImpl) GetMovieDetail(ImdbId string) (*entity.MovieDetail, error) {

	requestApiUrl := repo.requestApiUrl() + "&i=" + ImdbId

	movieDetail := &entity.MovieDetail{}

	response, err := http.Get(requestApiUrl)
	if err != nil {
		movieDetail.Response = "False"
		movieDetail.ErrorMessage = "Error when request api"
		return movieDetail, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(movieDetail)
	if err != nil {
		movieDetail.Response = "False"
		movieDetail.ErrorMessage = "Error when decode response"
		return movieDetail, err
	}

	return movieDetail, nil
}

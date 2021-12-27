package service

import (
	"context"
	"server-grpc/application/domain/entity"
	"server-grpc/application/services"
	"server-grpc/grpc/model"
)

type MovieServer struct {
	MovieServices *services.MovieServices
}

func (srv *MovieServer) Search(ctx context.Context, param *model.MovieSearchRequest) (*model.GetMovieListResponse, error) {
	listMovie, _ := srv.MovieServices.GetListMovie(&entity.MovieFilter{
		Search: param.Keyword,
		Page:   int(param.Page),
	})

	response := &model.GetMovieListResponse{
		TotalMovieResult: listMovie.TotalMovieResult,
		Response:         listMovie.Response,
		ErrorMessage:     listMovie.ErrorMessage,
	}

	for _, movie := range listMovie.Movies {
		response.Movies = append(response.Movies, &model.Movie{
			Title:  movie.Title,
			Imdbid: movie.ImdbId,
			Year:   movie.Year,
			Type:   movie.Type,
			Poster: movie.Poster,
		})
	}

	return response, nil
}

func (srv *MovieServer) Detail(ctx context.Context, param *model.GetMovieDetailRequest) (*model.GetMovieDetailResponse, error) {

	detailMovie, _ := srv.MovieServices.GetMovieDetail(param.ImdbID)

	response := &model.GetMovieDetailResponse{
		Title:        detailMovie.Title,
		Year:         detailMovie.Year,
		Rated:        detailMovie.Rated,
		Released:     detailMovie.Released,
		Runtime:      detailMovie.Runtime,
		Genre:        detailMovie.Genre,
		Director:     detailMovie.Director,
		Writer:       detailMovie.Writer,
		Actors:       detailMovie.Actors,
		Plot:         detailMovie.Plot,
		Language:     detailMovie.Language,
		Country:      detailMovie.Country,
		Awards:       detailMovie.Awards,
		Poster:       detailMovie.Poster,
		Metascore:    detailMovie.Metascore,
		ImdbRating:   detailMovie.ImdbRating,
		ImdbVotes:    detailMovie.ImdbVotes,
		ImdbID:       detailMovie.ImdbID,
		Type:         detailMovie.Type,
		DVD:          detailMovie.Dvd,
		BoxOffice:    detailMovie.BoxOffice,
		Production:   detailMovie.Production,
		Website:      detailMovie.Website,
		Response:     detailMovie.Response,
		ErrorMessage: detailMovie.ErrorMessage,
	}

	for _, rating := range detailMovie.Ratings {
		response.Ratings = append(response.Ratings, &model.Rating{
			Source: rating.Source,
			Value:  rating.Value,
		})
	}

	return response, nil
}

package services_test

import (
	"log"
	"server-grpc/application/domain/entity"
	"server-grpc/application/services"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func TestCreateLogSearch(t *testing.T) {
	movieService := services.SetupMovieServices()
	movies, err := movieService.GetListMovie(&entity.MovieFilter{
		Search: "Batman",
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(movies.Movies) == 0 {
		t.Fatal("no movie found")
	}

	if movies.Response != "True" {
		t.Fatal("response is not true")
	}
}

func TestGetMovieDetail(t *testing.T) {
	movieService := services.SetupMovieServices()
	movie, err := movieService.GetMovieDetail("tt0076759")

	if err != nil {
		t.Fatal(err)
	}

	if movie.Response != "True" {
		t.Fatal("response is not true")
	}

	if movie.ImdbID != "tt0076759" {
		t.Fatal("imdb id is not tt0076759")
	}
}

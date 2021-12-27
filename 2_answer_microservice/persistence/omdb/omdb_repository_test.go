package omdb_test

import (
	"log"
	"server-grpc/application/domain/entity"
	"server-grpc/config"
	"server-grpc/persistence/omdb"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func TestListMovie(t *testing.T) {

	omdbRepo := omdb.NewMovieOmdbRepository(
		config.OmdbApiKey(),
		config.OmdbApiUrl(),
	)

	result, err := omdbRepo.GetListMovie(&entity.MovieFilter{
		Search: "Batman",
	})

	if err != nil {
		t.Fatal(err)
	}

	if result.Response != "True" {
		t.Log(result.ErrorMessage)
		t.Fatal("Response imdb false")
	}

	t.Log("Total TotalMovieResult:", result.TotalMovieResult)

}

func TestGetDetailMovie(t *testing.T) {
	omdbRepo := omdb.NewMovieOmdbRepository(
		config.OmdbApiKey(),
		config.OmdbApiUrl(),
	)

	imdbID := "tt4853102"
	result, err := omdbRepo.GetMovieDetail(imdbID)

	if err != nil {
		t.Fatal(err)
	}

	if result.Response != "True" {
		t.Log(result.ErrorMessage)
		t.Fatal("Response imdb false")
	}

	if result.ImdbID != imdbID {
		t.Log("ImdbID:", result.ImdbID)
		t.Fatal("ImdbID not match")
	}

	if result.Title != "Batman: The Killing Joke" {
		t.Log("Title:", result.Title)
		t.Fatal("Title not match")
	}

	t.Log("Title:", result.Title)

}

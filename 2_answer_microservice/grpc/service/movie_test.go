package service_test

import (
	"context"
	"log"
	"net"
	"server-grpc/application/services"
	"server-grpc/grpc/model"
	"server-grpc/grpc/service"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var listener *bufconn.Listener

func init() {
	listener = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	model.RegisterMovieServicesServer(s, &service.MovieServer{
		MovieServices: services.SetupMovieServices(),
	})
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return listener.Dial()
}

func TestListMovie(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := model.NewMovieServicesClient(conn)
	response, err := client.Search(context.Background(), &model.MovieSearchRequest{
		Keyword: "Batman",
	})
	if err != nil {
		t.Fatal(err)
	}
	// check response length grt than 0
	if len(response.Movies) == 0 {
		t.Fatal("Response length is 0")
	}
}

func TestDetailMovie(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := model.NewMovieServicesClient(conn)
	response, err := client.Detail(context.Background(), &model.GetMovieDetailRequest{
		ImdbID: "tt0076759",
	})
	if err != nil {
		t.Fatal(err)
	}

	if response.ImdbID != "tt0076759" {
		t.Fatal("Response imdb id is not tt0076759")
	}

	if len(response.Ratings) == 0 {
		t.Fatal("Response Ratings length is 0")
	}
}

package handler

import (
	"log"
	"server-grpc/config"

	"server-grpc/grpc/model"
	"server-grpc/rest-api/service"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func ServiceMovieClient() model.MovieServicesClient {
	movieService := config.HostGrpc()
	conn, err := grpc.Dial(movieService, []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock()}...)

	if err != nil {
		log.Fatalf("could not connect to grpc server: %v", err)
	}

	return model.NewMovieServicesClient(conn)

}

func SetupRouters() *gin.Engine {
	r := gin.Default()
	serviceRestMovie := &service.ServiceRestMovie{
		MovieClient: ServiceMovieClient(),
	}
	r.GET("/v1/search-movie", serviceRestMovie.SearchMovie)
	r.GET("/v1/detail-movie/:imdbid", serviceRestMovie.DetailMovie)
	return r
}

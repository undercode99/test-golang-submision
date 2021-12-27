package service

import (
	"server-grpc/grpc/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServiceRestMovie struct {
	MovieClient model.MovieServicesClient
}

func (sr *ServiceRestMovie) SearchMovie(c *gin.Context) {
	keyword := c.Query("keyword")
	page := c.Query("page")
	number, err := strconv.ParseInt(page, 10, 32)
	if err != nil {
		number = 1
	}
	response, _ := sr.MovieClient.Search(c, &model.MovieSearchRequest{
		Keyword: keyword,
		Page:    int32(number),
	})

	c.JSON(200, gin.H{
		"Search": response,
	})
}

func (sr *ServiceRestMovie) DetailMovie(c *gin.Context) {
	imdbid := c.Param("imdbid")
	response, _ := sr.MovieClient.Detail(c, &model.GetMovieDetailRequest{
		ImdbID: imdbid,
	})
	c.JSON(200, response)
}

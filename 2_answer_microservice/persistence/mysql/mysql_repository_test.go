package mysql_test

import (
	"log"
	"server-grpc/config"
	"server-grpc/persistence/mysql"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func TestCreateLogSearch(t *testing.T) {
	db := mysql.NewMysqlConnection(config.MysqlDsn())
	movieMysqlRepo := mysql.NewMovieMysqlRepository(db)

	err := movieMysqlRepo.CreateLogSearch("Batman")

	if err != nil {
		t.Fatal(err)
	}
}

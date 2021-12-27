package config

import (
	"fmt"
	"os"

	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("WARNING: .env file not found")
	}

	return
}

func Env() string {
	if val, ok := os.LookupEnv("ENV"); ok {
		return val
	}

	return "development"
}

func MysqlDsn() string {
	dbUser := os.Getenv("MYSQL_DB_USER")
	dbPassword := os.Getenv("MYSQL_DB_PASSWORD")
	dbProtocol := os.Getenv("MYSQL_DB_PROTOCOL")
	dbHost := os.Getenv("MYSQL_DB_HOST")
	dbPort := os.Getenv("MYSQL_DB_PORT")
	dbName := os.Getenv("MYSQL_DB_NAME")

	return fmt.Sprintf("%v:%v@%v(%v:%v)/%v?parseTime=true",
		dbUser, dbPassword, dbProtocol, dbHost, dbPort, dbName)
}

func OmdbApiKey() string {
	val, ok := os.LookupEnv("OMDB_API_KEY")
	if !ok {
		log.Fatal("OMDB_API_KEY not provided")
	}

	return val
}

func OmdbApiUrl() string {
	val, ok := os.LookupEnv("OMDB_API_URL")
	if !ok {
		log.Fatal("OMDB_API_URL not provided")
	}

	return val
}

func GrpcPort() string {
	return fmt.Sprintf(":%s", os.Getenv("SERVER_GRPC_PORT"))
}

func RestPort() string {
	return fmt.Sprintf(":%s", os.Getenv("SERVER_REST_PORT"))
}

func HostGrpc() string {
	return os.Getenv("SERVER_GRPC_HOST") + GrpcPort()
}

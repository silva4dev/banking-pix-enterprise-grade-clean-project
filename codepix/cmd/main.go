package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/silva4dev/microservices-codepix-project/application/grpc"
	"github.com/silva4dev/microservices-codepix-project/infrastructure/db"
)

var database *gorm.DB

func main() {
	database := db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}

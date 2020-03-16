package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/edwintcloud/print-solution/server/api"
	"github.com/labstack/echo"

	mongo "github.com/edwintcloud/print-solution/server/repositories/mongo"
	"github.com/edwintcloud/print-solution/server/services/print"
	"github.com/globalsign/mgo"
)

const (
	serverAddr = "127.0.0.1:8081"
)

func main() {

	// connect mongo db
	session, err := connectMongoDB()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// create mongo repos
	dbName := os.Getenv("MONGO_DB")
	clientsRepo := mongo.NewMongoRepository(session, dbName, "clients")
	jobsRepo := mongo.NewMongoRepository(session, dbName, "jobs")

	// initialize print service
	printService := print.NewPrintService(clientsRepo, jobsRepo)

	// register api handlers and middlewares
	api := &api.API{
		PrintService: printService,
		E:            echo.New(),
	}
	api.RegisterHandlers()
	api.RegisterMiddlewares()

	// serve static content
	api.E.Static("/", "static")

	// start http server
	port := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	api.E.Logger.Fatal(api.E.Start(port))
}

func connectMongoDB() (*mgo.Session, error) {
	mongoURL := os.Getenv("MONGO_URL")
	timeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
	return mgo.DialWithTimeout(mongoURL, time.Duration(timeout)*time.Second)
}

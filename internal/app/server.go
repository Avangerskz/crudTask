package app

import (
	"context"
	"database/sql"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log"
	"net/http"
	"taskRestAPI/configs"
	"taskRestAPI/internal/app/handler"
	"taskRestAPI/internal/app/repository"
	service2 "taskRestAPI/internal/app/service"
	pb "taskRestAPI/proto"
)

//StartHTTPServer creates all dependencies ans starts api server on port 9090
func StartHTTPServer(ctx context.Context, cfg *configs.Configs, errCh chan error)  {
	log.Print("listening port:", cfg.Port)

	//router := mux.NewRouter()

	//open connection to postgres database on url from config.json
	db, dbErr := sql.Open("postgres", cfg.DbUrl)
	if dbErr != nil{
		log.Fatal(dbErr)
	}

	//creating CRUDRepository object with constructor that accepts sql.DB interface
	repo := repository.NewCRUDRepository(db)

	//calling a constructor that takes CRUDRepository interface and returns CRUDService object
	service := service2.NewCRUDService(repo)

	//creating api handler with constructor
	crudHandler := handler.NewCRUDHandler(service)

	//creating multiplexer for grpc-gateway
	muxServe := runtime.NewServeMux()

	//registering the http handlers for service
	if regErr := pb.RegisterCRUDServiceHandlerServer(ctx, muxServe, crudHandler); regErr!=nil{
		errCh <- regErr
	}

	//starting http server that takes port from configs and multiplexer
	err := http.ListenAndServe(cfg.Port, muxServe)
	if err != nil{
		log.Fatal("listen and serve error")
	}

}
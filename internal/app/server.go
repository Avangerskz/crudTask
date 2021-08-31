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

func StartHTTPServer(ctx context.Context, cfg *configs.Configs, errCh chan error)  {
	log.Print("listening port:", cfg.Port)

	//router := mux.NewRouter()

	db, dbErr := sql.Open("postgres", cfg.DbUrl)
	if dbErr != nil{
		log.Fatal(dbErr)
	}

	repo := repository.NewCRUDRepository(db)

	service := service2.NewCRUDService(repo)

	crudHandler := handler.NewCRUDHandler(service)

	muxServe := runtime.NewServeMux()
	if regErr := pb.RegisterCRUDServiceHandlerServer(ctx, muxServe, crudHandler); regErr!=nil{
		errCh <- regErr
	}

	err := http.ListenAndServe(cfg.Port, muxServe)
	if err != nil{
		log.Fatal("listen and serve error")
	}
}
package service

import (
	"context"
	"taskRestAPI/internal/app/repository"
	pb "taskRestAPI/proto"
)

type CRUDServiceImpl struct {
	repo repository.CRUDRepository
}

func (C CRUDServiceImpl) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	err := C.repo.CreateUser(ctx, req)
	if err != nil{
		return nil, err
	}
	res := &pb.CreateUserRes{}
	res.Msg = "user created"
	return res, nil
}

func (C CRUDServiceImpl) GetUserByUUID(ctx context.Context, req *pb.GetUserByUUIDReq) (*pb.GetUserByUUIDRes, error) {
	return C.repo.GetUserByUUID(ctx, req)
}

func (C CRUDServiceImpl) UpdateUserByUUID(ctx context.Context, req *pb.UpdateUserByUUIDReq) (*pb.UpdateUserByUUIDRes, error) {
	err := C.repo.UpdateUserByUUID(ctx, req)
	if err != nil{
		return nil, err
	}
	res := &pb.UpdateUserByUUIDRes{}
	res.Msg = "user updated"
	return res, nil
}

func NewCRUDService(repo repository.CRUDRepository) CRUDService {
	return &CRUDServiceImpl{repo: repo}
}
package repository

import (
	"context"
	"database/sql"
	"github.com/micro/go-micro/v2/errors"
	pb "taskRestAPI/proto"
)

type CRUDRepositoryImpl struct {
	db *sql.DB
}

func (c CRUDRepositoryImpl) CreateUser(ctx context.Context, req *pb.CreateUserReq)  error {
	tx, err := c.db.Begin()
	if err != nil{
		return errors.InternalServerError("internalserver", err.Error())
	}

	_, err = tx.ExecContext(
		ctx,
		"insert into users (firstname, lastname, email, age) values ($1, $2, $3, $4)",
		req.FirstName, req.LastName, req.Email, req.Age,
	)
	if err != nil{
		tx.Rollback()
		return errors.InternalServerError("internalserver", err.Error())
	}
	tx.Commit()
	return nil

}

func (c CRUDRepositoryImpl) GetUserByUUID(ctx context.Context, req *pb.GetUserByUUIDReq) (*pb.GetUserByUUIDRes, error) {
	tx, err := c.db.Begin()
	if err != nil{
		return nil, errors.InternalServerError("internalserver", err.Error())
	}

	res := &pb.GetUserByUUIDRes{}

	//var id int
	//rows, err := tx.QueryContext(
	//	ctx,
	//	"select * from users where id = $1",
	//	req.Uuid,
	//	)
	row := tx.QueryRowContext(
		ctx,
		"select * from users where id = $1",
		req.Uuid,
	)

	if row.Err() != nil{
		tx.Rollback()
		return nil, errors.InternalServerError("internalserver", row.Err().Error())
	}
	scanErr := row.Scan(&res.Uuid, &res.FirstName, &res.LastName, &res.Email, &res.Age, &res.CreatedDate)
	if scanErr != nil{
		tx.Rollback()
		return nil, errors.NotFound("notfound", scanErr.Error())
	}

	if res.FirstName == ""{
		tx.Rollback()
		return nil, errors.BadRequest("badrequest", "user not found")
	}

	tx.Commit()
	return res, nil
}

func (c CRUDRepositoryImpl) UpdateUserByUUID(ctx context.Context, req *pb.UpdateUserByUUIDReq) error {
	tx, err := c.db.Begin()
	if err != nil{
		return errors.InternalServerError("internalserver", err.Error())
	}

	_, err =  tx.QueryContext(
		ctx,
		"update users set firstname = $1, lastname = $2, email = $3, age = $4 where id = $5",
		req.FirstName, req.LastName, req.Email, req.Age, req.Uuid,
	)
	if err != nil{
		tx.Rollback()
		return errors.InternalServerError("internalserver", err.Error())
	}
	tx.Commit()
	return nil
}

func NewCRUDRepository(db *sql.DB) CRUDRepository {
	return &CRUDRepositoryImpl{db: db}
}

package models

import (
	"errors"
	"main/db"
	"main/types"
)

type UserModel interface {
	GetUsers(types.ReqUser) ([]types.User, error)
	GetUserByID(int) (types.User, error)
	CreateUser(types.ReqCreateUser) (types.User, error)
}

type userModel struct {
	db db.Database
}

func NewUserModel(db db.Database) UserModel {
	return &userModel{db}
}

func (um *userModel) GetUsers(req types.ReqUser) ([]types.User, error) {
	res := []types.User{}
	return res, errors.New("")
}

func (um *userModel) GetUserByID(id int) (types.User, error) {
	res := types.User{}
	return res, errors.New("")
}

func (um *userModel) CreateUser(req types.ReqCreateUser) (types.User, error) {
	res := types.User{}
	return res, errors.New("")
}

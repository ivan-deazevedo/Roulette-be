package repository

import "lunch-be/model"

type UserRepositoryInterface interface {
	InsertUser(model.PostUser) bool
	GetAllUsers() []model.User
	GetOneUser(uint) model.User
	DeleteUser(uint) bool
	UpdateUser(uint, model.UpdateUser) model.User
}

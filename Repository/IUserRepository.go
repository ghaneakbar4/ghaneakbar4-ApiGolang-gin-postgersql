package repository

import models "github.com/ghaneakbar4/ApiGolang/Models"
// اینترفیسی برای تمامی متدهایی که باید در ریپازیتوری پیاده سازی شوند
type IUserRepository interface {
	
	AddUser(models.User) bool
	GetAllUser() []models.User
	GetOneUser(uint) models.User
	UpdateUser(uint, models.User) models.User
	DeleteUser(uint) bool
}
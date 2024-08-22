package controller

import (
	"database/sql"

	models "github.com/ghaneakbar4/ApiGolang/Models"
	repository "github.com/ghaneakbar4/ApiGolang/Repository"
	"github.com/gin-gonic/gin"
)


type UserController struct {
	Db *sql.DB
}

// پیاده سازی اینترفیس ها توسط این کنترلر
func ImpUserController(db *sql.DB) UserControllerInterface {
	return &UserController{Db: db}
}

// حذف کاربر
func (m *UserController) DeleteUser(c *gin.Context) {
	DB := m.Db
	var uri models.UserUri
	//یک نمونه از ریپازیتوری
	repository := repository.ImpUserRepository(DB)
	delete := repository.DeleteUser(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "عملیات موفق"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "خطا"})
		return
	}
}

// گرفتن لیست تمامی کاربران
func (m *UserController) GetAllUsers(c *gin.Context) {
	DB := m.Db
	repository := repository.ImpUserRepository(DB)
	get := repository.GetAllUser()
	if (get != nil) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "عملیات موفق"})
		return
	} else {
		c.JSON(200, gin.H{"status": "success", "data": nil, "msg": "خطا"})
		return
	}
}

// دریافت اطلاعات یک کاربر با ای دی
func (m *UserController) GetOneUser(c *gin.Context) {
	DB := m.Db
	var uri models.UserUri

	repository := repository.ImpUserRepository(DB)
	get := repository.GetOneUser(uri.ID)
	if (get != models.User{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "عملیات موفق"})
		return
	} else {
		c.JSON(200, gin.H{"status": "success", "data": nil, "msg": "خطا"})
		return
	}
}

// متد ایجاد کاربر جدید
func (m *UserController) AddUser(c *gin.Context) {
	DB := m.Db
	var post models.User
	repository := repository.ImpUserRepository(DB)
	insert := repository.AddUser(post)
	if insert {
		c.JSON(200, gin.H{"status": "success", "msg": "کاربر جدید اضافه شد"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "خطا در افزودن  "})
		
	}
}

// 
func (m *UserController) UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{"status": "success", "msg": "jتمرین برای شما "})
	
}



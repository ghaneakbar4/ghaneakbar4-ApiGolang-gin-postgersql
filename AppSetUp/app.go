package appsetup

import (
	"database/sql"

	controller "github.com/ghaneakbar4/ApiGolang/Controller"
	db "github.com/ghaneakbar4/ApiGolang/DB"
	"github.com/gin-gonic/gin"
)

// ساختار دیتابیس و جین برای استفاده
type App struct {
	DB     *sql.DB
	Router *gin.Engine
}

// کانکتور دیتابیس
func (obj *App) DatabaseConection() {
	db := db.Connectdb()
	obj.DB = db
}

func (a *App) Routes() {
	r := gin.Default()
	controller := controller.ImpUserController(a.DB)
	r.POST("/api/AddUser", controller.AddUser)
	r.GET("/api/GetAllUsers", controller.GetAllUsers)
	r.GET("/api/GetOneUser/:id", controller.GetOneUser)
	r.PUT("/api/UpdateUser/:id", controller.UpdateUser)
	r.DELETE("/api/DeleteUser/:id", controller.DeleteUser)
	a.Router = r
}

// تابع راه اندازی سرور روی پورت 8000 
func (obj *App) Run() {
	obj.Router.Run(":8000")
}


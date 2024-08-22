package repository

import (
	"database/sql"
	"log"

	models "github.com/ghaneakbar4/ApiGolang/Models"
)

type UserRepository struct {
	Db *sql.DB
}
// پیاده سازی تمامی متدهای اینترفیس ریپازیتوری توسط این ریپازیتوری 
func ImpUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{Db: db}
}

// حذف کاربر
func (m *UserRepository) DeleteUser(id uint) bool {
	_, err := m.Db.Exec("DELETE FROM tbl_user WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		
		return false
	}
	return true
}

//دریافت لیست تمامی کاربران
func (m *UserRepository) GetAllUser() []models.User {
	query, err := m.Db.Query("SELECT * FROM tbl_user")
	if err != nil {
		log.Println(err)
		return nil
	}
	var mangas []models.User
	if query != nil {
		for query.Next() {
			var (
				id       uint
				firstname    string
				lastname    string
				username  string
				password string
				age   uint
				address string
			)
			err := query.Scan(&id, &firstname, &lastname, &username, &password, &age,&address)
			if err != nil {
				log.Println(err)
			}
			manga := models.User{FirstName:firstname,LastName: lastname,UserName: username,Password: password,Age: uint32(age),Address: address }
			mangas = append(mangas, manga)
		}
	}
	return mangas
}

// دریافت اطلاعات یک کاربر با ای دی
func (m *UserRepository) GetOneUser(id uint) models.User {
	query, err := m.Db.Query("SELECT * FROM tbl_user WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return models.User{}
	}
	var user models.User
	if query != nil {
		for query.Next() {
			var (
				id       uint
				firstname    string
				lastname    string
				username  string
				password string
				age   uint
				address string
			)
			err := query.Scan(&id, &firstname, &lastname, &username, &password, &age,&address)
			if err != nil {
				log.Println(err)
			}
			user = models.User{FirstName:firstname,LastName: lastname,UserName: username,Password: password,Age: uint32(age),Address: address }
		}
	}
	return user
}

// ریپازیتوری ایجاد یک کاربر جدید 
func (m *UserRepository) AddUser(post models.User) bool {
	stmt, err := m.Db.Prepare("INSERT INTO tbl_user(firstname, lastname, username, password, age,address) VALUES ($1,$2,$3,$4,$5,$6)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(post.FirstName, post.LastName, post.UserName, post.Password, post.Age,post.Address)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// ابدیت کردن یک کاربر
func (m *UserRepository) UpdateUser(id uint, post  models.User) models.User {
	//یک تمرین برای شما 
	return m.GetOneUser(id)
}
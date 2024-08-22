package db


/*
در این قسمت یک کانکتور برای ارتباط با دیتابیس ساخته شد 
 پکیج های اینجا 
 برای پستگرس و فریمورک جین میباشد 

*/
import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
   )
   // تابع اتصال به دیتابیس
   func Connectdb() *sql.DB {
	// کانکشن استرینگ و مهم نبودن اتصال ssl
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", UserDB, PASSDB, HOSTDB, DBNAME)
	// باز کردن دیتابیس
	db, err := sql.Open("postgres", connStr)
	//اگر خطایی پیش 
	if err != nil {
	 log.Fatal(err)
	}
	return db
   }
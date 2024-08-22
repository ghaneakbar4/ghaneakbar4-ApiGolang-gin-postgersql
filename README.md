## پروژه برای یادگیری دوستان علاقه مند به گولنگ این پروژه شامل عملیات crud بر روی دیتابیس پستگرس با استفاده از api و فریمورک جین میباشد
#مراحل راه اندازی پروژه
1-clone repository from github
2-get all packages by => go get . 
3-use command => go mod tidy
4-install postgesql and pgadmin on docker
5-change config.go values  in DB directory
6-use this commandes for run program => go run main.go
7- test by postman

//curl add
curl --location 'http://localhost:8000/api/AddUser' \
--header 'Content-Type: application/json' \
--data '{
    "firstname":"اکبر",
    "lastname":"رضاییان قانع",
    "username":"admin",
    "password":"1234",
    "age":100,
    "address":"تهران"

}'
//curl getall
curl --location 'http://localhost:8000/api/GetAllUsers'
//curl delete
curl --location --request DELETE 'http://localhost:8000/api/DeleteUser/2'




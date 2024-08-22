## پروژه برای یادگیری دوستان علاقه مند به گولنگ این پروژه شامل عملیات crud بر روی دیتابیس پستگرس با استفاده از api و فریمورک جین میباشد

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

curl --location 'http://localhost:8000/api/GetAllUsers'
curl --location --request DELETE 'http://localhost:8000/api/DeleteUser/2'


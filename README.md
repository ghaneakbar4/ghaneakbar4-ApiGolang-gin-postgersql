# پروژه برای یادگیری دوستان علاقه مند به گولنگ این پروژه شامل عملیات crud بر روی دیتابیس پستگرس با استفاده از api و فریمورک جین میباشد


## setup and run project step by step

## 1-clone repository from github
## 2-get all packages

Use the package manager 

```bash
go get .
```

## 3-use command 

```python
go mod tidy
```
## 4-install Postgersql and Pgadmin on Docker
## 5-change Config.go values in DB directory
## 6-use this Commandes for run program 
```
go run main.go
```
## 7- Test Api by postman

## Curl for postman
```
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
```
## script for generate user table 
```
-- Table: public.tbl_user

-- DROP TABLE IF EXISTS public.tbl_user;

CREATE TABLE IF NOT EXISTS public.tbl_user
(
    id bigint NOT NULL DEFAULT nextval('tbl_user_id_seq'::regclass),
    firstname character varying(255) COLLATE pg_catalog."default" NOT NULL,
    lastname character varying(255) COLLATE pg_catalog."default" NOT NULL,
    username character varying(255) COLLATE pg_catalog."default" NOT NULL,
    password character varying(255) COLLATE pg_catalog."default" NOT NULL,
    age integer,
    address character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT tbl_user_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.tbl_user
    OWNER to admin;
```


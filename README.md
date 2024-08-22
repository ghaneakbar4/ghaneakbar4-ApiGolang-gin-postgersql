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

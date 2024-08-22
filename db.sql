
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

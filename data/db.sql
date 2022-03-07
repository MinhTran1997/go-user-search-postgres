CREATE TABLE IF NOT EXISTS public.users
(
    id character varying(40) COLLATE pg_catalog."default" NOT NULL,
    username character varying(120) COLLATE pg_catalog."default",
    email character varying(120) COLLATE pg_catalog."default",
    phone character varying(45) COLLATE pg_catalog."default",
    date_of_birth date,
    CONSTRAINT users_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.users
    OWNER to postgres;
    
----------------------------------------------------------------

insert query:
INSERT INTO public.users VALUES ('omniman', 'omnimam.invicible', 'omni@gmail.com', '0523666888', '1965-11-16');

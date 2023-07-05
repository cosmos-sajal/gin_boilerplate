
-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    dob timestamp with time zone,
    mobile_number text,
    uuid uuid DEFAULT public.uuid_generate_v4()
);

ALTER TABLE public.users OWNER TO postgres;

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);

CREATE INDEX idx_users_uuid ON public.users USING btree (uuid);

-- +migrate Down

DROP SEQUENCE public.users_id_seq CASCADE;
DROP TABLE public.users CASCADE;

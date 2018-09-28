--
-- PostgreSQL database dump
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

CREATE TYPE public.book_type AS ENUM (
    'FICTION',
    'NONFICTION'
);

CREATE TABLE public.authors (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name text NOT NULL
);

CREATE TABLE public.books (
    id integer NOT NULL,
    author_id uuid NOT NULL,
    isbn character(32) NOT NULL,
    booktype public.book_type NOT NULL,
    title text NOT NULL,
    pages integer NOT NULL,
    summary text,
    available timestamp with time zone DEFAULT '2017-09-04 21:43:39.197538-04'::timestamp with time zone NOT NULL
);

CREATE SEQUENCE public.books_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE ONLY public.books ALTER COLUMN id SET DEFAULT nextval('public.books_id_seq'::regclass);


ALTER TABLE ONLY public.authors
    ADD CONSTRAINT authors_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_isbn_key UNIQUE (isbn);

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (id);

CREATE INDEX authors_name_idx ON public.authors USING btree (name);

CREATE INDEX books_title_idx ON public.books USING btree (author_id, title);

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.authors(id);

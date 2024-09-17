--
-- PostgreSQL database dump
--

-- Dumped from database version 15.7 (Debian 15.7-1.pgdg120+1)
-- Dumped by pg_dump version 16.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: todos; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.todos (
    id integer NOT NULL,
    title character varying(255) NOT NULL,
    description text,
    category character varying(100) NOT NULL,
    priority character varying(50) NOT NULL,
    is_completed boolean NOT NULL,
    is_removed boolean NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    user_id integer NOT NULL
);


--
-- Name: todos_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.todos_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: todos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.todos_id_seq OWNED BY public.todos.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: todos id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.todos ALTER COLUMN id SET DEFAULT nextval('public.todos_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: todos; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.todos (id, title, description, category, priority, is_completed, is_removed, created_at, updated_at, user_id) FROM stdin;
2	Buy groceries	Buy groceries for the week	home	medium	f	f	2024-08-01 00:00:00+00	2024-08-01 00:00:00+00	1
3	Workout	Workout for 1 hour	hobby	low	f	f	2024-08-01 00:00:00+00	2024-08-01 00:00:00+00	1
4	Work on the project 2	Work on the project for 2 hours	work	high	f	f	2024-08-01 00:00:00+00	2024-08-01 00:00:00+00	1
5	Buy groceries	Buy groceries for the week	home	medium	f	f	2024-08-01 00:00:00+00	2024-08-01 00:00:00+00	1
6	Workout	Workout for 1 hour	hobby	low	f	f	2024-08-01 00:00:00+00	2024-08-01 00:00:00+00	1
7	Work on the project 3	Work on the project for 2 hours	work	high	f	f	2024-08-01 00:00:00+00	2024-08-01 00:00:00+00	1
8	Buy groceries 2	Buy groceries for the week	home	medium	t	t	2024-08-01 00:00:00+00	2024-08-01 00:00:00+00	1
9	Work on the project 4	Work on the project for 2 hours	work	high	f	f	2024-08-01 00:00:00+00	2024-08-01 00:00:00+00	1
10	Buy groceries 3	Buy groceries for the week	home	medium	f	f	2024-08-01 00:00:00+00	2024-08-01 00:00:00+00	1
11	Workout 2	Workout for 1 hour	hobby	low	f	f	2024-08-01 00:00:00+00	2024-08-01 00:00:00+00	1
12	Work on the project 5	Work on the project for 2 hours	work	high	t	t	2024-08-01 00:00:00+00	2024-08-01 00:00:00+00	1
13	Work on the project 5	Work on the project for 2 hours	work	high	t	t	2024-08-01 00:00:00+00	2024-08-01 00:00:00+00	1
1	Work on the project 1	Work on the project for 2 hours	work	high	f	f	2024-08-01 00:00:00+00	2024-09-15 06:35:38.294015+00	1
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.users (id, username, email, password, created_at, updated_at) FROM stdin;
1	test1	test1@email.com	$2a$12$Gjf4qf84qmR1XARkkZqYq.1q/UlXy30JKN2ONIWbSIuYnuiuc2JvW	2024-08-01 00:00:00	2024-08-01 00:00:00
\.


--
-- Name: todos_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.todos_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- Name: todos todos_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.todos
    ADD CONSTRAINT todos_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: todos fk_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.todos
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--


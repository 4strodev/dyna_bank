--
-- PostgreSQL database dump
--

-- Dumped from database version 15.1
-- Dumped by pg_dump version 15.1

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
-- Name: accounts; Type: TABLE; Schema: application; Owner: dyna_bank_admin
--

CREATE TABLE application.accounts (
    id integer NOT NULL,
    name character varying(20) NOT NULL,
    user_id integer,
    description character varying(255),
    summary numeric(20,3) DEFAULT 0 NOT NULL
);


ALTER TABLE application.accounts OWNER TO dyna_bank_admin;

--
-- Name: accounts_id_seq; Type: SEQUENCE; Schema: application; Owner: dyna_bank_admin
--

CREATE SEQUENCE application.accounts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE application.accounts_id_seq OWNER TO dyna_bank_admin;

--
-- Name: accounts_id_seq; Type: SEQUENCE OWNED BY; Schema: application; Owner: dyna_bank_admin
--

ALTER SEQUENCE application.accounts_id_seq OWNED BY application.accounts.id;


--
-- Name: users; Type: TABLE; Schema: application; Owner: dyna_bank_admin
--

CREATE TABLE application.users (
    id integer NOT NULL,
    username character varying(20) NOT NULL,
    password character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    phone character varying(10) NOT NULL
);


ALTER TABLE application.users OWNER TO dyna_bank_admin;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: application; Owner: dyna_bank_admin
--

CREATE SEQUENCE application.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE application.users_id_seq OWNER TO dyna_bank_admin;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: application; Owner: dyna_bank_admin
--

ALTER SEQUENCE application.users_id_seq OWNED BY application.users.id;


--
-- Name: accounts id; Type: DEFAULT; Schema: application; Owner: dyna_bank_admin
--

ALTER TABLE ONLY application.accounts ALTER COLUMN id SET DEFAULT nextval('application.accounts_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: application; Owner: dyna_bank_admin
--

ALTER TABLE ONLY application.users ALTER COLUMN id SET DEFAULT nextval('application.users_id_seq'::regclass);


--
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: application; Owner: dyna_bank_admin
--

ALTER TABLE ONLY application.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: application; Owner: dyna_bank_admin
--

ALTER TABLE ONLY application.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: application; Owner: dyna_bank_admin
--

ALTER TABLE ONLY application.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: accounts fk_accounts_users; Type: FK CONSTRAINT; Schema: application; Owner: dyna_bank_admin
--

ALTER TABLE ONLY application.accounts
    ADD CONSTRAINT fk_accounts_users FOREIGN KEY (user_id) REFERENCES application.users(id);


--
-- Name: TABLE accounts; Type: ACL; Schema: application; Owner: dyna_bank_admin
--

GRANT SELECT,INSERT,DELETE,UPDATE ON TABLE application.accounts TO dyna_bank_application;


--
-- Name: SEQUENCE accounts_id_seq; Type: ACL; Schema: application; Owner: dyna_bank_admin
--

GRANT USAGE ON SEQUENCE application.accounts_id_seq TO dyna_bank_application;


--
-- Name: TABLE users; Type: ACL; Schema: application; Owner: dyna_bank_admin
--

GRANT SELECT,INSERT,DELETE,UPDATE ON TABLE application.users TO dyna_bank_application;


--
-- Name: SEQUENCE users_id_seq; Type: ACL; Schema: application; Owner: dyna_bank_admin
--

GRANT USAGE ON SEQUENCE application.users_id_seq TO dyna_bank_application;


--
-- PostgreSQL database dump complete
--


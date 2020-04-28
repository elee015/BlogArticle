CREATE DATABASE blogdb;
\connect blogdb;
CREATE TABLE articles ( id varchar(100) CONSTRAINT id_pk PRIMARY KEY, title varchar(80) NOT NULL, content varchar(250), author varchar(50) );


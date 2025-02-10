CREATE DATABASE golang_rss;
GRANT ALL PRIVILEGES ON DATABASE golang_rss TO postgres;

\connect golang_rss
CREATE SCHEMA IF NOT EXISTS dbo AUTHORIZATION postgres;
\q
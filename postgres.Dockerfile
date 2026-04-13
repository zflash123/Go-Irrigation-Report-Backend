FROM postgis/postgis:16-3.5

ARG DB_USER
ARG DB_PWD
ARG DB_NAME

ADD ./07-04-2026_postgresql-db-sql-backup-03.sql /docker-entrypoint-initdb.d/db-backup.sql

EXPOSE 5432
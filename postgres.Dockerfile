FROM postgis/postgis:16-master

ARG DB_USER
ARG DB_PWD
ARG DB_NAME

ENV POSTGRES_USER=${DB_USER}
ENV POSTGRES_PASSWORD=${DB_PWD}
ENV POSTGRES_NAME=${DB_NAME}

EXPOSE 5432

ADD ./07-04-2026_postgresql-db-sql-backup-03.sql /docker-entrypoint-initdb.d/db-backup.sql
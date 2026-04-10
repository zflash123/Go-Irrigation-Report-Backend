FROM postgis/postgis:17-master

ADD ./07-04-2026_postgresql-db-sql-backup-03.sql /docker-entrypoint-initdb.d/db-backup.sql
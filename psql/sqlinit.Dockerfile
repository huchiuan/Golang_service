FROM library/postgres
COPY /psql/init.sql /docker-entrypoint-initdb.d/
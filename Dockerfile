FROM postgres
COPY payment.sql /docker-entrypoint-initdb.d/
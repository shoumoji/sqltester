FROM mariadb:10.5

ENV MYSQL_DATABASE=testdb \
	MYSQL_ROOT_PASSWORD=root

COPY schema.sql /docker-entrypoint-initdb.d

EXPOSE 3306
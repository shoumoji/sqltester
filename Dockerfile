FROM mariadb:10.5

ENV MYSQL_DATABASE=testdb \
	MYSQL_ROOT_PASSWORD=root \
	character-set-server=utf8mb4

COPY schema.sql /docker-entrypoint-initdb.d
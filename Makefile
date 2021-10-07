# DB 起動待ち時間
BOOT_LATENCY_SEC := 10

DB_NAME := testdb

.PHONY: run
run:
	docker exec -it $$(docker run -d --name $(DB_NAME) --rm $$(docker build -q .)) \
		sh -c 'sleep $(BOOT_LATENCY_SEC) && mysql -u root -proot $(DB_NAME)'

.PHONY: stop
stop:
	docker stop $(DB_NAME)
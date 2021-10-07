.PHONY: run
run:
	docker run --rm -it $$(docker build -q .)
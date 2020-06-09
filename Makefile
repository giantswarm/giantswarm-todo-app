.PHONY: all docker-build docker-push dev-release

all: docker-build

docker-build: 
	cd api-server && $(MAKE) docker-build
	cd todo-manager && $(MAKE) docker-build

docker-push:
	cd api-server && $(MAKE) docker-push
	cd todo-manager && $(MAKE) docker-push

dev-release: docker-build
	kubectl -n todo rollout restart deployment apiserver
	kubectl -n todo rollout restart deployment todomanager



install:
	go get ./...
	go mod vendor

build-frontend-img:
	@echo "Building frontend docker image"
	docker build -t frontend -f ./frontend/Dockerfile .

build-stripe-event-handler-img:
	@echo "Building stripe-events-handler docker image"
	docker build -t stripe-events-handler -f ./stripe-events-handler/Dockerfile .

build-imgs: build-frontend-img build-stripe-event-handler-img


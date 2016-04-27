NGINX_VERSION:=1.9.15
IMAGE_NAME=robinmonjo/nginx-module:dev

build:
	docker build --build-arg NGINX_VERSION=$(NGINX_VERSION) -t $(IMAGE_NAME) .

test: build
	docker run -w /lab/integration/ $(IMAGE_NAME) go test

clean:
	docker rmi -f $(IMAGE_NAME)
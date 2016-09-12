NGINX_VERSION:=1.10.1
NDK_VERSION:=0.3.0
GO_VERSION:=1.7.1
IMAGE_NAME:=robinmonjo/nginx-module:dev

build:
	docker build --build-arg NGINX_VERSION=$(NGINX_VERSION) --build-arg NDK_VERSION=$(NDK_VERSION) --build-arg GO_VERSION=$(GO_VERSION) -t $(IMAGE_NAME) .

test: build
	docker run -w /lab/src/integration/ -e "DOMAIN=test.io" $(IMAGE_NAME) go test

clean:
	docker rmi -f $(IMAGE_NAME)

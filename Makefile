PACKAGE_PATH = github.com\/SaucySalamander
PROJECT_NAME = owl-db
VERSION = 0.0.1-alpha

export PATH := $(PATH):$(HOME)/go/bin

owl-db:
	go build -o ./$(PROJECT_NAME)-$(VERSION)

image: image-build image-clean

image-build:
	sed -e "s/PACKAGE_PATH/$(PACKAGE_PATH)/g" -e "s/PROJECT_NAME/$(PROJECT_NAME)/g" -e "s/VERSION_NUMBER/$(VERSION)/g" ./build/Dockerfile > Dockerfile
	docker build . -f ./Dockerfile --tag $(PROJECT_NAME):$(VERSION)

image-clean:
	rm ./Dockerfile
	
chart:

publish:

proto:
	protoc -I. -I/usr/local/include/ --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/proto/*.proto

clean:
	go clean
	rm ./Dockerfile
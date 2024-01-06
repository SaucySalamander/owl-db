PACKAGE_PATH = github.com\/SaucySalamander
PROJECT_NAME = owl-db
VERSION = 0.0.1-alpha

image: image-build image-clean

image-build:
	sed -e "s/PACKAGE_PATH/$(PACKAGE_PATH)/g" -e "s/PROJECT_NAME/$(PROJECT_NAME)/g" -e "s/VERSION_NUMBER/$(VERSION)/g" ./build/Dockerfile > Dockerfile
	docker build . -f ./Dockerfile --tag $(PROJECT_NAME):$(VERSION)

image-clean:
	rm ./Dockerfile
	
chart:

publish:

clean:
	go clean
	rm ./Dockerfile
PACKAGE_PATH = github.com/SaucySalamander
PROJECT_NAME = owl-db
VERSION = 0.0.1-alpha

owl-db:
	go build -o ./$(PROJECT_NAME)-$(VERSION)

image:

chart:

publish:

clean:
	go clean
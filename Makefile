.PHONY: build

default: build

build:
	    go install
	    GOOS=darwin GOARCH=amd64 go build -o build/darwin/terraform-provider-gha github.com/patrickmarabeas/terraform-provider-gha
	    GOOS=linux GOARCH=amd64 go build -o build/linux/terraform-provider-gha github.com/patrickmarabeas/terraform-provider-gha

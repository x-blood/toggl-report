.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/main handlers/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose --aws-profile $(PROFILE)

deploy_prd: clean build
	sls deploy --verbose --stage prd --aws-profile $(PROFILE)
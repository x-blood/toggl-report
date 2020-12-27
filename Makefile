.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/main handlers/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose --aws-profile $(PROFILE)

deploy_prd: clean build
	sls deploy --verbose --stage prd --aws-profile $(PROFILE)
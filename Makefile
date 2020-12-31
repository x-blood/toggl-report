.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/main handlers/main.go

clean:
	rm -rf ./bin

deploy:
	sls deploy --verbose --aws-profile $(PROFILE)

deploy_prd:
	sls deploy --verbose --stage prd --aws-profile $(PROFILE)
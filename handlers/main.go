package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"toggl-report/controller"
)

func Handler() error {
	return controller.Do()
}

func main() {
	lambda.Start(Handler)
}

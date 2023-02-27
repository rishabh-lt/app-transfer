package main

import (
	"github.com/LambdatestIncPrivate/app-transfer/controller"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(controller.TransferHandler)
}

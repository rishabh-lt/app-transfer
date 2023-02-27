package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LambdatestIncPrivate/app-transfer/cloud"
	"github.com/LambdatestIncPrivate/app-transfer/logger"
	"github.com/LambdatestIncPrivate/app-transfer/models"
	"github.com/aws/aws-lambda-go/events"
)

func TransferHandler(apiRequest events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := logger.InitLogging()

	var body models.DataTransferRequest

	err := json.Unmarshal([]byte(apiRequest.Body), &body)

	if err != nil {
		errMsg := fmt.Sprintf("Unable to parse JSON. Error : " + err.Error())
		logger.Error(errMsg)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       errMsg,
		}, nil
	}

	copyOutput, err := cloud.Transfer(body, logger)

	if err != nil {
		errMsg := "unable to copy file from source to destination"
		logger.Error(errMsg)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       errMsg,
		}, nil
	}
	logger.Info(copyOutput.GoString())
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "success",
	}, nil
}

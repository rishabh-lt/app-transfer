package cloud

import (
	"fmt"

	"github.com/LambdatestIncPrivate/app-transfer/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"go.uber.org/zap"
)

func Transfer(transferRequest models.DataTransferRequest, logger *zap.Logger) (*s3.CopyObjectOutput, error) {
	destinationSession := getSession(transferRequest.DestinationRegion)

	s3Conn := s3.New(destinationSession)
	input := &s3.CopyObjectInput{
		Bucket:     aws.String(transferRequest.DestinationBucket),
		CopySource: aws.String(fmt.Sprintf("%v/%v", transferRequest.SourceBucket, transferRequest.SourcePath)),
		Key:        aws.String(transferRequest.SourcePath),
	}

	result, err := s3Conn.CopyObject(input)
	if err != nil {
		errMsg := fmt.Sprintf("Couldn't copy object from %v:%v to %v:%v. Here's why: %v\n",
			transferRequest.SourceBucket, transferRequest.SourcePath, transferRequest.DestinationBucket, transferRequest.SourcePath, err.Error())
		logger.Error(errMsg)
		return nil, err
	}

	return result, nil
}

func getSession(region string) *session.Session {

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	return sess

}

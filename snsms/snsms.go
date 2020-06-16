package snsms

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func GetClient(accessKey string, secretKey string, region string) (*sns.SNS, error) {
	credentials := credentials.NewStaticCredentials(accessKey, secretKey, "")
	session, err := session.NewSession(&aws.Config{
		Credentials: credentials,
		Region:      aws.String(region),
	})
	if err != nil {
		return nil, err
	}

	return sns.New(session), nil
}

// 送信メッセージ、送信対象の電話番号を引数にとり、PublishInputのインスタンスを作る
func CreateInputMessage(msg string, phoneNum string) *sns.PublishInput {
	return &sns.PublishInput{
		Message:     aws.String(msg),
		PhoneNumber: aws.String(phoneNum),
	}
}

// subscriberにむけてのPublishInputを作る
func CreateInputMessageToSubscriber(msg string, topicArn string) *sns.PublishInput {
	return &sns.PublishInput{
		Message:  aws.String(msg),
		TopicArn: aws.String(topicArn),
	}
}

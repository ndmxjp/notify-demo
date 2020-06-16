package ses

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	// The character encoding for the email.
	CharSet = "UTF-8"
)

func NewClient(accessKey, secretKey, region string) (*ses.SES, error) {
	credentials := credentials.NewStaticCredentials(accessKey, secretKey, "")
	session, err := session.NewSession(&aws.Config{
		Credentials: credentials,
		Region:      aws.String(region),
	})
	if err != nil {
		return nil, err
	}
	return ses.New(session), nil
}

func CreateInputMessage(from, to, sub, body string) *ses.SendEmailInput {
	return &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(to),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(sub),
			},
		},
		Source: aws.String(from),
	}
}

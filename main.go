package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"firebase.google.com/go/messaging"
	"github.com/ndmxjp/notify-demo/firebase"
	"github.com/ndmxjp/notify-demo/ses"
	snssms "github.com/ndmxjp/notify-demo/snsms"
)

const (
	awsAccessKey   string = "AWS_ACCESS_KEY"
	awsSecret      string = "AWS_SECRET_KEY"
	region         string = "ap-northeast-1"
	sesRegion      string = "us-west-2"
	credentialPath string = "GOOGLE_APPLICATION_CREDENTIALS"
)

const (
	sender   string = "SENDER"
	to       string = "TO"
	phoneNum string = "PHONE_NUM"
	fcmToken string = "FCM_TOKEN"
)

func main() {
	awsAK := os.Getenv(awsAccessKey)
	awsSec := os.Getenv(awsSecret)
	credentialPath := os.Getenv(credentialPath)
	snsClient, err := snssms.GetClient(awsAK, awsSec, region)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	sesClient, err := ses.NewClient(awsAK, awsSec, sesRegion)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	firebaseClient, err := firebase.NewClient(credentialPath)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	// ses mail 送信
	input := ses.CreateInputMessage(os.Getenv(sender), os.Getenv(to), "test ses from code", "test ses from code")
	result, err := sesClient.SendEmail(input)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	fmt.Printf("Result: %s", result.String())

	// sms送信
	// 必須国コード
	msgIn := snssms.CreateInputMessage("TestMessage", os.Getenv(phoneNum))

	smsResult, err := snsClient.Publish(msgIn)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	fmt.Printf("Result: %s", smsResult.String())

	// push通知
	ctx := context.Background()
	messageClient, _ := firebaseClient.Messaging(ctx)
	// This registration token comes from the client FCM SDKs.
	registrationToken := os.Getenv(fcmToken)

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := messageClient.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
}

package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"fmt"
	"google.golang.org/api/option"
	"log"
)

var (
	client *messaging.Client
	ctx    = context.Background()
)

func main() {
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}
	client, err = app.Messaging(ctx)
	if err != nil {
		log.Fatalf("getting Messaging client error: %v", err)
	}

}

func sendNotification(title string, body string, fcmToken string) {
	// This registration token comes from the client FCM SDKs.
	registrationToken := fcmToken

	notification := &messaging.Notification{
		Title: title,
		Body:  body,
	}

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"Type":    title,
			"Content": body,
		},
		Token:        registrationToken,
		Notification: notification,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		fmt.Println("[sendNotification] client.Send error :", err, registrationToken)
	}
	_ = response
}

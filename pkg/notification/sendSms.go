package notification

import (
	"go-ecommerce-app/configs"
)

type NotificationClient interface {
	SendSMS(phone string, msg string) error
}

type notificationClient struct {
	config configs.AppConfig
}

func NewNotificationClient(config configs.AppConfig) NotificationClient {
	return &notificationClient{
		config: config,
	}
}

//Twilio
func (c *notificationClient) SendSMS(phone string, msg string) error {
	
	return nil
}

package notification

import (
	"encoding/json"
	"fmt"
	"go-ecommerce-app/configs"
	"strings"

	"github.com/twilio/twilio-go"

	voice "github.com/twilio/twilio-go/rest/api/v2010"
)

type NotificationClient interface {
	SendVoiceCall(phone string, msg string) error
}

type notificationClient struct {
	config configs.AppConfig
}

func NewNotificationClient(config configs.AppConfig) NotificationClient {
	return &notificationClient{
		config: config,
	}
}

// Twilio
func (c *notificationClient) SendVoiceCall(phone string, code string) error {
	accountSID := c.config.AccountSID
	authToken := c.config.AuthToken

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSID,
		Password: authToken,
	})

	spacedCode := strings.Join(strings.Split(code, ""), " ")
	message := fmt.Sprintf("Your verification code for ecommerce site is %s", spacedCode)

	//Use TwiML Bin URL or dynamic response
	twiml := fmt.Sprintf("<Response><Say>%s</Say></Response>", message)

	params := &voice.CreateCallParams{}
	params.SetTo(phone)
	params.SetFrom(c.config.TwilioPhoneNo)
	params.SetTwiml(twiml)

	resp, err := client.Api.CreateCall(params)
	if err != nil {
		fmt.Println("Error making voice call:" + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("voice call response:" + string(response))
	}

	return err
}

package email

import (
	"fmt"
	"github.com/mailjet/mailjet-apiv3-go"
	"log"
)

func SendEmail(recipient string, variablesMap map[string]interface{}) error {
	m := mailjet.NewMailjetClient(MJ_APIKEY_PUBLIC, MJ_APIKEY_PRIVATE)
	messageInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: SENDER_EMAIL,
				Name:  "Manu :)",
			},
			To: &mailjet.RecipientsV31{
				{
					Email: recipient,
					Name:  recipient,
				},
			},
			TemplateID:       TEMPLATE_ID,
			TemplateLanguage: true,
			Subject:          "Test challenge!",
			Variables:        variablesMap,
		},
	}
	messages := mailjet.MessagesV31{Info: messageInfo}
	res, err := m.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("Data: %+v\n", res)
	return nil
}

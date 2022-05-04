package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"

	"github.com/pkg/errors"
)

const (
	telegramMessageTemplate = `{{ .Icon }} {{ .Title }}
{{ .Message }}
`
)

// telegramNotifier defines structure of notifier specific to Telegram
type TelegramNotifier struct {
	Notifier

	// endpoint holds information about telegram endpoint
	endpoint string
	// token holds information about bot token used to send notifications
	token string
	// targets holds list of targets who should receive notification
	targets []string
}

// NewTelegramNotifier creates new instance of Telegram Notifier
func NewTelegramNotifier(token string) *TelegramNotifier {
	return &TelegramNotifier{
		endpoint: "https://api.telegram.org",
		token:    token,
	}
}

// DeliverNotification delivers notification using logic for specific notifier
func (notifier TelegramNotifier) DeliverNotification(notification *Notification, targets []string) error {
	notifier.targets = targets
	return notifier.deliverNotification(notification)
}

// telegramRequest represents structure of request sent to Telegram API
type telegramRequest struct {
	ChatID    string `json:"chat_id"`
	Photo     string `json:"photo,omitempty"`
	Caption   string `json:"caption,omitempty"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

// telegramResponse represents structure of Telegram API response
type telegramResponse struct {
	Success     bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

// deliverNotification is an internal method which delivers notification using logic for specific notifier
func (notifier TelegramNotifier) deliverNotification(notification *Notification) error {
	message, err := notifier.messageBuilder(notification)

	if err != nil {
		return err
	}

	for _, target := range notifier.targets {
		var rawPayload telegramRequest
		var sendMethod string
		if notification.HasImage() {
			rawPayload = telegramRequest{
				ChatID:    target,
				Photo:     notification.Image(),
				Caption:   message,
				ParseMode: "Markdown",
			}
			sendMethod = "sendPhoto"
		} else {
			rawPayload = telegramRequest{
				ChatID:    target,
				Text:      message,
				ParseMode: "Markdown",
			}
			sendMethod = "sendMessage"
		}

		payload, err := json.Marshal(rawPayload)
		if err != nil {
			return err
		}

		rawResponse, err := sendPostRequest(
			fmt.Sprintf(
				"%s/bot%s/%s",
				notifier.endpoint,
				notifier.token,
				sendMethod,
			),
			payload,
		)

		if err != nil {
			return err
		}

		response := telegramResponse{}
		err = json.Unmarshal(rawResponse, &response)
		if err != nil {
			return err
		}

		if !response.Success {
			return errors.Errorf("%d: %s", response.ErrorCode, response.Description)
		}
	}

	return nil
}

// messageBuilder builds message from provided notification
func (notifier TelegramNotifier) messageBuilder(notification *Notification) (string, error) {
	tpl, err := template.New("").Parse(telegramMessageTemplate)
	if err != nil {
		return "", err
	}
	value := struct {
		Icon    string
		Title   string
		Message string
	}{
		Icon:    notification.Icon(),
		Title:   notification.Title(),
		Message: notification.Message(),
	}

	var templateBuffer bytes.Buffer
	err = tpl.Execute(&templateBuffer, value)
	if err != nil {
		return "", err
	}
	return templateBuffer.String(), nil
}

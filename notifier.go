package notifier

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Notifier knows how to deliver notification
type Notifier interface {
	// deliverNotification is an internal method which delivers notification using logic for specific notifier
	deliverNotification(notification *Notification) error
}

// sendPostRequest sends post request with provided data to specified endpoint
func sendPostRequest(endpoint string, data []byte) ([]byte, error) {
	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	return body, nil
}

package groupme

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/kfcampbell/pigskin/responses"
)

const baseURL = "https://api.groupme.com/v3"

// PostMessage posts a message to GroupMe
func PostMessage(message string, groupID string, apiKey string) error {
	url := baseURL + "/groups/" + groupID + "/messages?token=" + apiKey
	messageID := uuid.New()

	innerPacket := &responses.InnerMessage{
		SourceGUID: messageID.String(),
		Text:       message,
	}

	packet := &responses.Message{
		Message: *innerPacket,
	}

	body, err := json.Marshal(packet)
	if err != nil {
		return err
	}
	fmt.Printf("request body: %v\n", string(body))

	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	fmt.Printf("received %v posting to GroupMe\n", res.StatusCode)

	bodyBytes, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		return err
	}

	fmt.Printf("body: %v\n", string(bodyBytes))
	return nil
}

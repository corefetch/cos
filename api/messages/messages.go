package messages

import (
	"bytes"
	"encoding/json"
	"errors"
	"gom/core/service"
	"gom/pod/messages/ob"
	"net/http"
)

type EmailPayload struct {
	Email    string   `json:"email"`
	Template string   `json:"template"`
	Args     *ob.Args `json:"vars"`
}

func Send(to ob.Recipient, template string, args *ob.Args) (err error) {

	if args == nil {
		args = &ob.Args{}
	}

	(*args)["FIRST_NAME"] = to.SendName()[0]

	payload := EmailPayload{
		Email:    to.SendTo(),
		Template: template,
		Args:     args,
	}

	buf := &bytes.Buffer{}

	if err := json.NewEncoder(buf).Encode(payload); err != nil {
		return err
	}

	request, err := service.EndpointRequest("messages", "POST", "/send", buf)

	if err != nil {
		return err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New("status code is not 200")
	}

	return nil
}

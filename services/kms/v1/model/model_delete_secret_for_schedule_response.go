package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteSecretForScheduleResponse struct {
	Secret         *Secret `json:"secret,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteSecretForScheduleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSecretForScheduleResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecretForScheduleResponse", string(data)}, " ")
}

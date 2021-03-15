package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteSubscriptionResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteSubscriptionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"DeleteSubscriptionResponse", string(data)}, " ")
}

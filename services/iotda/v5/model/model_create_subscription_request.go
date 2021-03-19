package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateSubscriptionRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	Body *CreateSubReq `json:"body,omitempty"`
}

func (o CreateSubscriptionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionRequest", string(data)}, " ")
}

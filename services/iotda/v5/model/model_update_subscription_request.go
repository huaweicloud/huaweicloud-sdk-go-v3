package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateSubscriptionRequest struct {
	InstanceId     *string       `json:"Instance-Id,omitempty"`
	SubscriptionId string        `json:"subscription_id"`
	Body           *UpdateSubReq `json:"body,omitempty"`
}

func (o UpdateSubscriptionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionRequest", string(data)}, " ")
}

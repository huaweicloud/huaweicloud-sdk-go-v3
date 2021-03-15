package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteSubscriptionRequest struct {
	InstanceId     *string `json:"Instance-Id,omitempty"`
	SubscriptionId string  `json:"subscription_id"`
}

func (o DeleteSubscriptionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"DeleteSubscriptionRequest", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowSubscriptionRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	SubscriptionId string `json:"subscription_id"`
}

func (o ShowSubscriptionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"ShowSubscriptionRequest", string(data)}, " ")
}

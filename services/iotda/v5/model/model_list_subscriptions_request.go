package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListSubscriptionsRequest struct {
	InstanceId  *string `json:"Instance-Id,omitempty"`
	Resource    *string `json:"resource,omitempty"`
	Event       *string `json:"event,omitempty"`
	Callbackurl *string `json:"callbackurl,omitempty"`
	AppId       *string `json:"app_id,omitempty"`
	Channel     *string `json:"channel,omitempty"`
	Limit       *int32  `json:"limit,omitempty"`
	Marker      *string `json:"marker,omitempty"`
	Offset      *int32  `json:"offset,omitempty"`
}

func (o ListSubscriptionsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSubscriptionsRequest struct{}"
	}

	return strings.Join([]string{"ListSubscriptionsRequest", string(data)}, " ")
}

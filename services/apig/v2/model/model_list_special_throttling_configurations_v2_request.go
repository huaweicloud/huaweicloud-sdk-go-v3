package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListSpecialThrottlingConfigurationsV2Request struct {
	InstanceId string `json:"instance_id"`

	ThrottleId string `json:"throttle_id"`

	ObjectType *string `json:"object_type,omitempty"`

	AppName *string `json:"app_name,omitempty"`

	Offset *int64 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSpecialThrottlingConfigurationsV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSpecialThrottlingConfigurationsV2Request struct{}"
	}

	return strings.Join([]string{"ListSpecialThrottlingConfigurationsV2Request", string(data)}, " ")
}

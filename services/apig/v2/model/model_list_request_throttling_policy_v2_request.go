package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListRequestThrottlingPolicyV2Request struct {
	InstanceId    string  `json:"instance_id"`
	Id            *string `json:"id,omitempty"`
	Name          *string `json:"name,omitempty"`
	Offset        *int64  `json:"offset,omitempty"`
	Limit         *int32  `json:"limit,omitempty"`
	PreciseSearch *string `json:"precise_search,omitempty"`
}

func (o ListRequestThrottlingPolicyV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListRequestThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"ListRequestThrottlingPolicyV2Request", string(data)}, " ")
}

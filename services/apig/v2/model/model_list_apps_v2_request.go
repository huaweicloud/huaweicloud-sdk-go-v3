package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAppsV2Request struct {
	InstanceId string `json:"instance_id"`

	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *int32 `json:"status,omitempty"`

	AppKey *string `json:"app_key,omitempty"`

	Creator *string `json:"creator,omitempty"`

	Offset *int64 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	PreciseSearch *string `json:"precise_search,omitempty"`
}

func (o ListAppsV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAppsV2Request struct{}"
	}

	return strings.Join([]string{"ListAppsV2Request", string(data)}, " ")
}

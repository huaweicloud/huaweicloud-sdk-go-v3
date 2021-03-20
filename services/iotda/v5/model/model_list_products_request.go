package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListProductsRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	AppId *string `json:"app_id,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListProductsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProductsRequest struct{}"
	}

	return strings.Join([]string{"ListProductsRequest", string(data)}, " ")
}

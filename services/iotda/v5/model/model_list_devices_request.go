package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListDevicesRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	ProductId *string `json:"product_id,omitempty"`

	GatewayId *string `json:"gateway_id,omitempty"`

	IsCascadeQuery *bool `json:"is_cascade_query,omitempty"`

	NodeId *string `json:"node_id,omitempty"`

	DeviceName *string `json:"device_name,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	AppId *string `json:"app_id,omitempty"`
}

func (o ListDevicesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListDevicesRequest struct{}"
	}

	return strings.Join([]string{"ListDevicesRequest", string(data)}, " ")
}

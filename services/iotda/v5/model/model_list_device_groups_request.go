package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListDeviceGroupsRequest struct {
	InstanceId       *string `json:"Instance-Id,omitempty"`
	Limit            *int32  `json:"limit,omitempty"`
	Marker           *string `json:"marker,omitempty"`
	Offset           *int32  `json:"offset,omitempty"`
	LastModifiedTime *string `json:"last_modified_time,omitempty"`
	AppId            *string `json:"app_id,omitempty"`
}

func (o ListDeviceGroupsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListDeviceGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListDeviceGroupsRequest", string(data)}, " ")
}

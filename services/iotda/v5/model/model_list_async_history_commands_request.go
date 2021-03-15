package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAsyncHistoryCommandsRequest struct {
	DeviceId    string  `json:"device_id"`
	InstanceId  *string `json:"Instance-Id,omitempty"`
	Limit       *int32  `json:"limit,omitempty"`
	Marker      *string `json:"marker,omitempty"`
	Offset      *int32  `json:"offset,omitempty"`
	StartTime   *string `json:"start_time,omitempty"`
	EndTime     *string `json:"end_time,omitempty"`
	Status      *string `json:"status,omitempty"`
	CommandId   *string `json:"command_id,omitempty"`
	CommandName *string `json:"command_name,omitempty"`
}

func (o ListAsyncHistoryCommandsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAsyncHistoryCommandsRequest struct{}"
	}

	return strings.Join([]string{"ListAsyncHistoryCommandsRequest", string(data)}, " ")
}

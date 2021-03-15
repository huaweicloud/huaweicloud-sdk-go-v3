package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowPartitionMessageRequest struct {
	InstanceId    string `json:"instance_id"`
	Topic         string `json:"topic"`
	Partition     int32  `json:"partition"`
	MessageOffset string `json:"message_offset"`
}

func (o ShowPartitionMessageRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPartitionMessageRequest struct{}"
	}

	return strings.Join([]string{"ShowPartitionMessageRequest", string(data)}, " ")
}

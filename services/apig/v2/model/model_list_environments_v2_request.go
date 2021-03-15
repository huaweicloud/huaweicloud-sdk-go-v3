package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListEnvironmentsV2Request struct {
	InstanceId string  `json:"instance_id"`
	Name       *string `json:"name,omitempty"`
	Offset     *int64  `json:"offset,omitempty"`
	Limit      *int32  `json:"limit,omitempty"`
}

func (o ListEnvironmentsV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListEnvironmentsV2Request struct{}"
	}

	return strings.Join([]string{"ListEnvironmentsV2Request", string(data)}, " ")
}

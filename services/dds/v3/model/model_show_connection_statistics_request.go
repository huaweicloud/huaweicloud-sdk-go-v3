package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowConnectionStatisticsRequest struct {
	InstanceId string `json:"instance_id"`

	NodeId *string `json:"node_id,omitempty"`
}

func (o ShowConnectionStatisticsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowConnectionStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowConnectionStatisticsRequest", string(data)}, " ")
}

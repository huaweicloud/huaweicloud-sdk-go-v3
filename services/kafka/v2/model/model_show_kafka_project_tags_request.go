package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowKafkaProjectTagsRequest struct {
}

func (o ShowKafkaProjectTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowKafkaProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowKafkaProjectTagsRequest", string(data)}, " ")
}

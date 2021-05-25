package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowRabbitMqProjectTagsRequest struct {
}

func (o ShowRabbitMqProjectTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowRabbitMqProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowRabbitMqProjectTagsRequest", string(data)}, " ")
}

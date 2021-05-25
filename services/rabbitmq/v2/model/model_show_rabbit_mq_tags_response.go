package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowRabbitMqTagsResponse struct {
	// 标签列表

	Tags           *[]CreateInstanceReqTags `json:"tags,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowRabbitMqTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowRabbitMqTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowRabbitMqTagsResponse", string(data)}, " ")
}

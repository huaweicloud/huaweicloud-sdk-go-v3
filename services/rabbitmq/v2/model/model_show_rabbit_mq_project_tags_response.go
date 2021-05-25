package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowRabbitMqProjectTagsResponse struct {
	// 标签列表

	Tags           *[]ShowProjectTagsRespTags `json:"tags,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowRabbitMqProjectTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowRabbitMqProjectTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowRabbitMqProjectTagsResponse", string(data)}, " ")
}

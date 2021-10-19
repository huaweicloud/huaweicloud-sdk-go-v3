package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowKafkaProjectTagsResponse struct {
	// 标签列表

	Tags           *[]TagMultyValueEntity `json:"tags,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowKafkaProjectTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowKafkaProjectTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowKafkaProjectTagsResponse", string(data)}, " ")
}

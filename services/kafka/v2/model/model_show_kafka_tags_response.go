package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowKafkaTagsResponse struct {
	// 标签列表

	Tags           *[]CreatePostPaidInstanceReqTags `json:"tags,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ShowKafkaTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowKafkaTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowKafkaTagsResponse", string(data)}, " ")
}

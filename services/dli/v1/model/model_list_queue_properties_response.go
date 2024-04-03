package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueuePropertiesResponse Response Object
type ListQueuePropertiesResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	Properties     *[]QueueProperty `json:"properties,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListQueuePropertiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueuePropertiesResponse struct{}"
	}

	return strings.Join([]string{"ListQueuePropertiesResponse", string(data)}, " ")
}

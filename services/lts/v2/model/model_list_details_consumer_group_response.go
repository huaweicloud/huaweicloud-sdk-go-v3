package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDetailsConsumerGroupResponse Response Object
type ListDetailsConsumerGroupResponse struct {
	Body           *[]ConsumerCheckpointInfo `json:"body,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListDetailsConsumerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDetailsConsumerGroupResponse struct{}"
	}

	return strings.Join([]string{"ListDetailsConsumerGroupResponse", string(data)}, " ")
}

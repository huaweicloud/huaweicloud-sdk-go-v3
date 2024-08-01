package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConsumerGroupResponse Response Object
type ListConsumerGroupResponse struct {
	Body           *[]ConsumerGroupInfo `json:"body,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListConsumerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConsumerGroupResponse struct{}"
	}

	return strings.Join([]string{"ListConsumerGroupResponse", string(data)}, " ")
}

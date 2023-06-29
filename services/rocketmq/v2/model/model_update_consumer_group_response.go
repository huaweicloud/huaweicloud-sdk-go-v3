package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConsumerGroupResponse Response Object
type UpdateConsumerGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateConsumerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConsumerGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateConsumerGroupResponse", string(data)}, " ")
}

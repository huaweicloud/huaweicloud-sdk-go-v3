package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateConsumerGroupResponse struct {
	// 消费组信息。

	Groups         *[]CreateConsumerGroupRespGroups `json:"groups,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o CreateConsumerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConsumerGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateConsumerGroupResponse", string(data)}, " ")
}

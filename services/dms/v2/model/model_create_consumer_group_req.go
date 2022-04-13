package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateConsumerGroupReq struct {
	// 消费组信息。  每个队列最多能创建3个消费组，如果请求中的消费组个数超过3个，请求校验不通过，无法创建消费组。

	Groups []GroupEntity `json:"groups"`
}

func (o CreateConsumerGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConsumerGroupReq struct{}"
	}

	return strings.Join([]string{"CreateConsumerGroupReq", string(data)}, " ")
}

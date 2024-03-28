package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateConsumerGroupReq struct {

	// 消费组列表。
	Groups *[]CreateOrUpdateConsumerGroup `json:"groups,omitempty"`
}

func (o BatchUpdateConsumerGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateConsumerGroupReq struct{}"
	}

	return strings.Join([]string{"BatchUpdateConsumerGroupReq", string(data)}, " ")
}

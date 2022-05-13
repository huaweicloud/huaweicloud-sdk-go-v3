package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteConsumerGroupReq struct {

	// 待删除的消费组列表。
	Groups *[]string `json:"groups,omitempty"`
}

func (o BatchDeleteConsumerGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteConsumerGroupReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteConsumerGroupReq", string(data)}, " ")
}

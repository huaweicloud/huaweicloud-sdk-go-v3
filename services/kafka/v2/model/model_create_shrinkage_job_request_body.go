package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateShrinkageJobRequestBody struct {

	// 缩容后集群节点数
	NewBrokerNum *string `json:"new_broker_num,omitempty"`
}

func (o CreateShrinkageJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShrinkageJobRequestBody struct{}"
	}

	return strings.Join([]string{"CreateShrinkageJobRequestBody", string(data)}, " ")
}

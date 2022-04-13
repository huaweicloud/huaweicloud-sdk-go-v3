package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PriorityBody struct {
	// 副本优先级，取值范围是0到100，0为默认禁止倒换。

	SlavePriorityWeight int32 `json:"slave_priority_weight"`
}

func (o PriorityBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PriorityBody struct{}"
	}

	return strings.Join([]string{"PriorityBody", string(data)}, " ")
}

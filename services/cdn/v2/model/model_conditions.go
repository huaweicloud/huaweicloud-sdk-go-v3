package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Conditions **参数解释：** 触发当前规则的条件 **约束限制：** 不涉及
type Conditions struct {
	Match *Match `json:"match"`
}

func (o Conditions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Conditions struct{}"
	}

	return strings.Join([]string{"Conditions", string(data)}, " ")
}

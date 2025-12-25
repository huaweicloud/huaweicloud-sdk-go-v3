package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyEnabled **参数解释** 是否启用告警策略 **约束限制** 不涉及 **取值范围** - true:开启 - false:关闭 **默认取值** true
type PolicyEnabled struct {
}

func (o PolicyEnabled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyEnabled struct{}"
	}

	return strings.Join([]string{"PolicyEnabled", string(data)}, " ")
}

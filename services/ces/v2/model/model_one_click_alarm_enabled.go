package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OneClickAlarmEnabled **参数解释** 是否启用一键告警 **约束限制** 不涉及 **取值范围** - true:开启 - false：关闭 **默认取值** true
type OneClickAlarmEnabled struct {
}

func (o OneClickAlarmEnabled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OneClickAlarmEnabled struct{}"
	}

	return strings.Join([]string{"OneClickAlarmEnabled", string(data)}, " ")
}

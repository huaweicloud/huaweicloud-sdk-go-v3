package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OneClickAlarmDescription **参数解释** 一键告警描述 **约束限制** 不涉及 **取值范围** 字符长度在0到256之间 **默认取值** 空字符串
type OneClickAlarmDescription struct {
}

func (o OneClickAlarmDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OneClickAlarmDescription struct{}"
	}

	return strings.Join([]string{"OneClickAlarmDescription", string(data)}, " ")
}

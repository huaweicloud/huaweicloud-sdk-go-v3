package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OneClickAlarmId **参数解释**： 一键告警ID。 **约束限制**： 不涉及。 **取值范围**： 只能为字母或者数字，字符长度为[1,64] **默认取值**： 不涉及。
type OneClickAlarmId struct {
}

func (o OneClickAlarmId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OneClickAlarmId struct{}"
	}

	return strings.Join([]string{"OneClickAlarmId", string(data)}, " ")
}

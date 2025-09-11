package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OneClickAlarmIdResp **参数解释**： 一键告警ID。 **取值范围**： 只能为字母或者数字，字符长度为[1,64]
type OneClickAlarmIdResp struct {
}

func (o OneClickAlarmIdResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OneClickAlarmIdResp struct{}"
	}

	return strings.Join([]string{"OneClickAlarmIdResp", string(data)}, " ")
}

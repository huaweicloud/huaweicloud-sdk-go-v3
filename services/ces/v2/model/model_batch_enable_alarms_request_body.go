package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchEnableAlarmsRequestBody struct {

	// 需要批量启停的告警规则的ID列表
	AlarmIds []string `json:"alarm_ids"`

	// 是否开启告警规则。true:开启，false:关闭。
	AlarmEnabled bool `json:"alarm_enabled"`
}

func (o BatchEnableAlarmsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchEnableAlarmsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchEnableAlarmsRequestBody", string(data)}, " ")
}

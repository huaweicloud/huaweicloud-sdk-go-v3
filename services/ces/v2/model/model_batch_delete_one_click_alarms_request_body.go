package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteOneClickAlarmsRequestBody struct {

	// 需要批量删除的一键告警ID列表
	OneClickAlarmIds []string `json:"one_click_alarm_ids"`
}

func (o BatchDeleteOneClickAlarmsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteOneClickAlarmsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteOneClickAlarmsRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteOneClickAlarmsRequestBody struct {

	// **参数解释**： 需要批量删除的一键告警ID列表。 **约束限制**： 一键告警ID数量最多为100个，至少1个。
	OneClickAlarmIds []string `json:"one_click_alarm_ids"`
}

func (o BatchDeleteOneClickAlarmsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteOneClickAlarmsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteOneClickAlarmsRequestBody", string(data)}, " ")
}

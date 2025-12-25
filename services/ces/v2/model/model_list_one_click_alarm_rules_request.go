package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOneClickAlarmRulesRequest Request Object
type ListOneClickAlarmRulesRequest struct {

	// **参数解释** 一键告警ID **约束限制** 不涉及 **取值范围** 长度为1到64字符，只能包含字母数字 **默认取值** 不涉及
	OneClickAlarmId string `json:"one_click_alarm_id"`
}

func (o ListOneClickAlarmRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOneClickAlarmRulesRequest struct{}"
	}

	return strings.Join([]string{"ListOneClickAlarmRulesRequest", string(data)}, " ")
}

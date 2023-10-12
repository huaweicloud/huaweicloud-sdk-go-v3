package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOneClickAlarmRulesRequest Request Object
type ListOneClickAlarmRulesRequest struct {

	// 一键告警ID
	OneClickAlarmId string `json:"one_click_alarm_id"`
}

func (o ListOneClickAlarmRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOneClickAlarmRulesRequest struct{}"
	}

	return strings.Join([]string{"ListOneClickAlarmRulesRequest", string(data)}, " ")
}

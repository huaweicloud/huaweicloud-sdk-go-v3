package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchEnableAlarmRulesResponse Response Object
type BatchEnableAlarmRulesResponse struct {

	// **参数解释**： 成功启停的告警规则ID列表。
	AlarmIds       *[]string `json:"alarm_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchEnableAlarmRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchEnableAlarmRulesResponse struct{}"
	}

	return strings.Join([]string{"BatchEnableAlarmRulesResponse", string(data)}, " ")
}

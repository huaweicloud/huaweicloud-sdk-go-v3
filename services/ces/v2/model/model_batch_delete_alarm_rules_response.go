package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteAlarmRulesResponse struct {

	// 成功删除的告警规则ID列表
	AlarmIds       *[]string `json:"alarm_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteAlarmRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAlarmRulesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteAlarmRulesResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateAlarmRulesResponse struct {

	// 告警规则id，以al开头，包含22个数字或字母
	AlarmId        *string `json:"alarm_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAlarmRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlarmRulesResponse struct{}"
	}

	return strings.Join([]string{"CreateAlarmRulesResponse", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmRulesResponse Response Object
type ListAlarmRulesResponse struct {

	// 告警规则列表
	Alarms *[]ListAlarmResponseAlarms `json:"alarms,omitempty"`

	// 告警规则总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAlarmRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmRulesResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmRulesResponse", string(data)}, " ")
}

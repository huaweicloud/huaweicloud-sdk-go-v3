package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAlarmRulesResponse struct {

	// 告警规则列表
	Alarms *[]ListAlarmResponseAlarms `json:"alarms,omitempty" xml:"alarms"`

	// 告警规则总数
	Count          *int32 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAlarmRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmRulesResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmRulesResponse", string(data)}, " ")
}

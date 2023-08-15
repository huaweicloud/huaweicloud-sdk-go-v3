package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmRuleResourcesRequest Request Object
type ListAlarmRuleResourcesRequest struct {

	// Alarm实例ID
	AlarmId string `json:"alarm_id"`

	// 分页偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAlarmRuleResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmRuleResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmRuleResourcesRequest", string(data)}, " ")
}

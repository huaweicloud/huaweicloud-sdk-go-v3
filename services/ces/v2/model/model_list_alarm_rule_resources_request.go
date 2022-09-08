package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAlarmRuleResourcesRequest struct {

	// 发送的实体的MIME类型。默认使用application/json; charset=UTF-8。
	ContentType string `json:"Content-Type"`

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

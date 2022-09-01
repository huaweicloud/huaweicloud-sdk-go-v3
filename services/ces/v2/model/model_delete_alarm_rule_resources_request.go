package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAlarmRuleResourcesRequest struct {

	// 发送的实体的MIME类型。默认使用application/json; charset=UTF-8。
	ContentType string `json:"Content-Type" xml:"Content-Type"`

	// Alarm实例ID
	AlarmId string `json:"alarm_id" xml:"alarm_id"`

	Body *ResourcesReqV2 `json:"body,omitempty" xml:"body"`
}

func (o DeleteAlarmRuleResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmRuleResourcesRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlarmRuleResourcesRequest", string(data)}, " ")
}

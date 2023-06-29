package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlarmRuleResourcesRequest Request Object
type DeleteAlarmRuleResourcesRequest struct {

	// Alarm实例ID
	AlarmId string `json:"alarm_id"`

	Body *ResourcesReqV2 `json:"body,omitempty"`
}

func (o DeleteAlarmRuleResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmRuleResourcesRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlarmRuleResourcesRequest", string(data)}, " ")
}

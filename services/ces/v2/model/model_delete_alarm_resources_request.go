package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAlarmResourcesRequest struct {
	// Alarm实例ID

	AlarmId string `json:"alarm_id"`

	Body *ResourcesReqV2 `json:"body,omitempty"`
}

func (o DeleteAlarmResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmResourcesRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlarmResourcesRequest", string(data)}, " ")
}

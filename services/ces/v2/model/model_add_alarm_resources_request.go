package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddAlarmResourcesRequest struct {
	// Alarm实例ID

	AlarmId string `json:"alarm_id"`

	Body *ResourcesReqV2 `json:"body,omitempty"`
}

func (o AddAlarmResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAlarmResourcesRequest struct{}"
	}

	return strings.Join([]string{"AddAlarmResourcesRequest", string(data)}, " ")
}

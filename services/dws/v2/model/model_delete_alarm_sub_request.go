package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAlarmSubRequest struct {

	// 告警订阅ID
	AlarmSubId string `json:"alarm_sub_id"`
}

func (o DeleteAlarmSubRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmSubRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlarmSubRequest", string(data)}, " ")
}

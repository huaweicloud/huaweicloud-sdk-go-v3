package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmInfo struct {

	// 告警设备
	AlarmDevice *string `json:"alarm_device,omitempty"`

	// 告警数量
	AlarmNumber *int32 `json:"alarm_number,omitempty"`

	// 告警时间
	AlarmDuration *int64 `json:"alarm_duration,omitempty"`
}

func (o AlarmInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmInfo struct{}"
	}

	return strings.Join([]string{"AlarmInfo", string(data)}, " ")
}

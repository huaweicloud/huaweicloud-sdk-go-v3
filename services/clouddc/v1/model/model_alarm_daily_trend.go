package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmDailyTrend struct {

	// 告警时间
	AlarmDate *sdktime.SdkTime `json:"alarm_date,omitempty"`

	// 告警设备信息，包括设备类型及告警数量
	AlarmDevices *[]AlarmDeviceInfo `json:"alarm_devices,omitempty"`

	// 告警级别
	AlarmLevels *[]AlarmLevelInfo `json:"alarm_levels,omitempty"`
}

func (o AlarmDailyTrend) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmDailyTrend struct{}"
	}

	return strings.Join([]string{"AlarmDailyTrend", string(data)}, " ")
}

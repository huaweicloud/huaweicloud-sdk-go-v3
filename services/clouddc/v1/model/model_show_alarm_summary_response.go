package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlarmSummaryResponse Response Object
type ShowAlarmSummaryResponse struct {

	// 告警级别
	AlarmLevels *[]AlarmLevelInfo `json:"alarm_levels,omitempty"`

	// 告警设备信息，包括设备类型及告警数量
	AlarmDevices *[]AlarmDeviceInfo `json:"alarm_devices,omitempty"`

	// 故障服务器Top10
	AlarmHosts *[]AlarmHost `json:"alarm_hosts,omitempty"`

	// 告警分组
	AlarmGroups    *[]AlarmGroup `json:"alarm_groups,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowAlarmSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmSummaryResponse struct{}"
	}

	return strings.Join([]string{"ShowAlarmSummaryResponse", string(data)}, " ")
}

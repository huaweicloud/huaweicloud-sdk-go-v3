package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmGroup struct {

	// 标签
	Label *string `json:"label,omitempty"`

	// 告警设备
	AlarmDevices *[]AlarmDeviceInfo `json:"alarm_devices,omitempty"`

	// sn列表
	Sns *[]string `json:"sns,omitempty"`
}

func (o AlarmGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmGroup struct{}"
	}

	return strings.Join([]string{"AlarmGroup", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmDeviceInfo struct {

	// 告警硬件
	AlarmDevice *string `json:"alarm_device,omitempty"`

	// 告警数量
	AlarmNumber *int32 `json:"alarm_number,omitempty"`
}

func (o AlarmDeviceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmDeviceInfo struct{}"
	}

	return strings.Join([]string{"AlarmDeviceInfo", string(data)}, " ")
}

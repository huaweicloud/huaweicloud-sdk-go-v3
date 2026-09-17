package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IoTdbPushInfo 创建IoTDB推送信息
type IoTdbPushInfo struct {
	DeviceData *DeviceIoTdbPushInfo `json:"device_data,omitempty"`
}

func (o IoTdbPushInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IoTdbPushInfo struct{}"
	}

	return strings.Join([]string{"IoTdbPushInfo", string(data)}, " ")
}

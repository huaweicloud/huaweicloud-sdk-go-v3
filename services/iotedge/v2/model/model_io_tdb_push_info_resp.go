package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IoTdbPushInfoResp IoTDB推送信息详情
type IoTdbPushInfoResp struct {
	DeviceData *DeviceIoTdbPushInfoDetail `json:"device_data,omitempty"`
}

func (o IoTdbPushInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IoTdbPushInfoResp struct{}"
	}

	return strings.Join([]string{"IoTdbPushInfoResp", string(data)}, " ")
}

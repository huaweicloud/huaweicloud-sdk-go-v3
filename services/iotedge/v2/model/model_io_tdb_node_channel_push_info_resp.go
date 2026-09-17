package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IoTdbNodeChannelPushInfoResp IoTDB推送信息详情
type IoTdbNodeChannelPushInfoResp struct {
	DeviceData *DeviceIoTdbNodeChannelPushInfoDetail `json:"device_data,omitempty"`
}

func (o IoTdbNodeChannelPushInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IoTdbNodeChannelPushInfoResp struct{}"
	}

	return strings.Join([]string{"IoTdbNodeChannelPushInfoResp", string(data)}, " ")
}

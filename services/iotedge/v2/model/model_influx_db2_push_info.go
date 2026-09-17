package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InfluxDb2PushInfo 创建IoTDB推送信息
type InfluxDb2PushInfo struct {
	DeviceData *DeviceInfluxDb2PushInfo `json:"device_data,omitempty"`
}

func (o InfluxDb2PushInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InfluxDb2PushInfo struct{}"
	}

	return strings.Join([]string{"InfluxDb2PushInfo", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InfluxDb2NodeChannelPushInfoRsp 创建IoTDB推送信息
type InfluxDb2NodeChannelPushInfoRsp struct {
	DeviceData *DeviceInfluxDb2NodeChannelPushInfoDetail `json:"device_data,omitempty"`
}

func (o InfluxDb2NodeChannelPushInfoRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InfluxDb2NodeChannelPushInfoRsp struct{}"
	}

	return strings.Join([]string{"InfluxDb2NodeChannelPushInfoRsp", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InfluxDb2NodeChannelDetailDto MQTT通道配置详情
type InfluxDb2NodeChannelDetailDto struct {
	ConnectionInfo *InfluxDb2ConnectionInfo `json:"connection_info"`

	PushInfo *InfluxDb2NodeChannelPushInfoRsp `json:"push_info"`
}

func (o InfluxDb2NodeChannelDetailDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InfluxDb2NodeChannelDetailDto struct{}"
	}

	return strings.Join([]string{"InfluxDb2NodeChannelDetailDto", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInfluxDb2ChannelDetail MQTT通道配置详情
type CreateInfluxDb2ChannelDetail struct {
	ConnectionInfo *InfluxDb2ConnectionInfo `json:"connection_info"`

	PushInfo *InfluxDb2PushInfo `json:"push_info"`
}

func (o CreateInfluxDb2ChannelDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInfluxDb2ChannelDetail struct{}"
	}

	return strings.Join([]string{"CreateInfluxDb2ChannelDetail", string(data)}, " ")
}

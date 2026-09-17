package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIoTdbChannelDetail MQTT通道配置详情
type CreateIoTdbChannelDetail struct {
	ConnectionInfo *IoTdbConnectionInfo `json:"connection_info"`

	PushInfo *IoTdbPushInfo `json:"push_info"`
}

func (o CreateIoTdbChannelDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIoTdbChannelDetail struct{}"
	}

	return strings.Join([]string{"CreateIoTdbChannelDetail", string(data)}, " ")
}

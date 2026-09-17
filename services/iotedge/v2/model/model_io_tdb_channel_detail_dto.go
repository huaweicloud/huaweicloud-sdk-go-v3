package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IoTdbChannelDetailDto IoTDB通道详情
type IoTdbChannelDetailDto struct {
	ConnectionInfo *IoTdbConnectionInfoResp `json:"connection_info,omitempty"`

	PushInfo *IoTdbPushInfoResp `json:"push_info,omitempty"`
}

func (o IoTdbChannelDetailDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IoTdbChannelDetailDto struct{}"
	}

	return strings.Join([]string{"IoTdbChannelDetailDto", string(data)}, " ")
}

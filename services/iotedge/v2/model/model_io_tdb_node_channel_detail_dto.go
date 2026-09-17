package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IoTdbNodeChannelDetailDto IoTDB通道详情
type IoTdbNodeChannelDetailDto struct {
	ConnectionInfo *IoTdbConnectionInfoResp `json:"connection_info,omitempty"`

	PushInfo *IoTdbNodeChannelPushInfoResp `json:"push_info,omitempty"`
}

func (o IoTdbNodeChannelDetailDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IoTdbNodeChannelDetailDto struct{}"
	}

	return strings.Join([]string{"IoTdbNodeChannelDetailDto", string(data)}, " ")
}

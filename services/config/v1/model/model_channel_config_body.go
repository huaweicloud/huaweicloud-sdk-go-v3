package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChannelConfigBody tracker通道配置
type ChannelConfigBody struct {
	Smn *TrackerSmnChannelConfigBody `json:"smn,omitempty"`

	Obs *TrackerObsChannelConfigBody `json:"obs,omitempty"`
}

func (o ChannelConfigBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChannelConfigBody struct{}"
	}

	return strings.Join([]string{"ChannelConfigBody", string(data)}, " ")
}

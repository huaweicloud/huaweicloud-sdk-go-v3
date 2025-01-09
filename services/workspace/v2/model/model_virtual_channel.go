package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VirtualChannel struct {

	// 是否开启虚拟通道策略设置。取值为：false：表示关闭。true：表示开启。
	VirtualChannelControlEnable *bool `json:"virtual_channel_control_enable,omitempty"`

	Options *VirtualChannelOptions `json:"options,omitempty"`
}

func (o VirtualChannel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VirtualChannel struct{}"
	}

	return strings.Join([]string{"VirtualChannel", string(data)}, " ")
}

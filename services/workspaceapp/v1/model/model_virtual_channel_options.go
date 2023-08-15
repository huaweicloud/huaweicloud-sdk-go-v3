package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VirtualChannelOptions struct {

	// 自定义虚拟通道注册名。
	CustomVirtualChannelName *string `json:"custom_virtual_channel_name,omitempty"`

	// 虚拟通道下载配置信息，需Base64加密。
	VirtualChannelPluginDetails *string `json:"virtual_channel_plugin_details,omitempty"`

	// 第三方插件名称。
	ThirdPartyPluginName *string `json:"third_party_plugin_name,omitempty"`
}

func (o VirtualChannelOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VirtualChannelOptions struct{}"
	}

	return strings.Join([]string{"VirtualChannelOptions", string(data)}, " ")
}

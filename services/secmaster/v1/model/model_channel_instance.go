package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChannelInstance struct {

	// 分组名称
	ChannelName *string `json:"channel_name,omitempty"`

	ConfigStatus *ConfigStatus `json:"config_status,omitempty"`

	// Iam用户ID
	CreateBy *string `json:"create_by,omitempty"`

	// 是否在线
	MiniOnOnline *bool `json:"mini_on_online,omitempty"`

	Monitor *Monitor `json:"monitor,omitempty"`

	// 分组名称
	NodeName *string `json:"node_name,omitempty"`

	// IP地址
	PrivateIpAddress *string `json:"private_ip_address,omitempty"`

	// IP地址
	PublicIpAddress *string `json:"public_ip_address,omitempty"`

	ReadWrite *ReadWrite `json:"read_write,omitempty"`

	// 地区
	Region *string `json:"region,omitempty"`
}

func (o ChannelInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChannelInstance struct{}"
	}

	return strings.Join([]string{"ChannelInstance", string(data)}, " ")
}

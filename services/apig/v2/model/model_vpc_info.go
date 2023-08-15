package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VpcInfo VPC通道详情。如果vpc_channel_status = 1，则这个object类型为必填信息
type VpcInfo struct {

	// 云服务器ID
	EcsId *string `json:"ecs_id,omitempty"`

	// 云服务器名称
	EcsName *string `json:"ecs_name,omitempty"`

	// 是否使用级联方式  暂不支持
	CascadeFlag *bool `json:"cascade_flag,omitempty"`

	// 代理主机
	VpcChannelProxyHost *string `json:"vpc_channel_proxy_host,omitempty"`

	// VPC通道编号
	VpcChannelId *string `json:"vpc_channel_id,omitempty"`

	// VPC通道端口
	VpcChannelPort *int32 `json:"vpc_channel_port,omitempty"`
}

func (o VpcInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcInfo struct{}"
	}

	return strings.Join([]string{"VpcInfo", string(data)}, " ")
}

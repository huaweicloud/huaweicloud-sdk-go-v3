package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VPC通道详情。如果vpc_channel_status = 1，则这个object类型为必填信息
type VpcInfo struct {

	// 云服务器ID
	EcsId *string `json:"ecs_id,omitempty" xml:"ecs_id"`

	// 云服务器名称
	EcsName *string `json:"ecs_name,omitempty" xml:"ecs_name"`

	// 是否使用级联方式  暂不支持
	CascadeFlag *bool `json:"cascade_flag,omitempty" xml:"cascade_flag"`

	// 代理主机
	VpcChannelProxyHost *string `json:"vpc_channel_proxy_host,omitempty" xml:"vpc_channel_proxy_host"`

	// VPC通道编号
	VpcChannelId *string `json:"vpc_channel_id,omitempty" xml:"vpc_channel_id"`

	// VPC通道端口
	VpcChannelPort *int32 `json:"vpc_channel_port,omitempty" xml:"vpc_channel_port"`
}

func (o VpcInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcInfo struct{}"
	}

	return strings.Join([]string{"VpcInfo", string(data)}, " ")
}

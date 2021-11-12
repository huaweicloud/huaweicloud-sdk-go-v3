package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VPC通道详情。vpc_channel_status = 1，则这个object类型为必填信息
type ApiBackendVpcReq struct {
	// 代理主机

	VpcChannelProxyHost *string `json:"vpc_channel_proxy_host,omitempty"`
	// VPC通道编号

	VpcChannelId string `json:"vpc_channel_id"`
}

func (o ApiBackendVpcReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiBackendVpcReq struct{}"
	}

	return strings.Join([]string{"ApiBackendVpcReq", string(data)}, " ")
}

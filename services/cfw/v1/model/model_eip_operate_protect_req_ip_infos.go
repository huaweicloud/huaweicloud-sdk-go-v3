package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EipOperateProtectReqIpInfos struct {

	// 弹性公网IP数据ID
	Id *string `json:"id,omitempty"`

	// 弹性公网IP地址
	PublicIp *string `json:"public_ip,omitempty"`

	// 弹性公网IP地址IPV6
	PublicIpv6 *string `json:"public_ipv6,omitempty"`
}

func (o EipOperateProtectReqIpInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipOperateProtectReqIpInfos struct{}"
	}

	return strings.Join([]string{"EipOperateProtectReqIpInfos", string(data)}, " ")
}

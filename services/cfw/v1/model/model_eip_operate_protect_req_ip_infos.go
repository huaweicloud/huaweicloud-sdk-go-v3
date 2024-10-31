package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EipOperateProtectReqIpInfos struct {

	// 弹性公网IP ID，可通过调用弹性IP列表查询接口获得，通过返回值中的data.records.id（.表示各对象之间层级的区分）获得。
	Id *string `json:"id,omitempty"`

	// 弹性公网IP IPv4地址，可通过调用弹性IP列表查询接口获得，通过返回值中的data.records.public_ip（.表示各对象之间层级的区分）获得。
	PublicIp *string `json:"public_ip,omitempty"`

	// 弹性公网IP IPv6地址，可通过调用弹性IP列表查询接口获得，通过返回值中的data.records.public_ipv6（.表示各对象之间层级的区分）获得。
	PublicIpv6 *string `json:"public_ipv6,omitempty"`
}

func (o EipOperateProtectReqIpInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipOperateProtectReqIpInfos struct{}"
	}

	return strings.Join([]string{"EipOperateProtectReqIpInfos", string(data)}, " ")
}

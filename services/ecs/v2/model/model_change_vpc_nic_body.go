package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeVpcNicBody This is a auto create Body Object
type ChangeVpcNicBody struct {

	// 网卡ID，UUID格式。 当该字段不为空时，表示挂载指定的网卡。port_id和subnet_id不能同时为空。 网卡ID可以从虚拟私有云的“查询端口列表”章节查询到。 约束： 网卡必须带有安全组。 网卡状态必须为DOWN。 网卡的vpcid必须和传入的vpcid一致。 当port_id和subnet_id同时存在的时候，优先使用port_id。当选择port_id不为空时，代表此时使用的是弹性网卡，此时security_groups和ip_address等参数不生效。
	PortId *string `json:"port_id,omitempty"`

	// 云服务器云主机添加网卡的信息。 需要指定云服务器云主机所属虚拟私有云下已创建的网络（network）的ID，UUID格式。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 添加网卡的安全组信息
	SecurityGroups *[]ChangeVpcSecurityGroups `json:"security_groups,omitempty"`

	// P地址，无该参数表示自动分配IP地址
	IpAddress *string `json:"ip_address,omitempty"`
}

func (o ChangeVpcNicBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeVpcNicBody struct{}"
	}

	return strings.Join([]string{"ChangeVpcNicBody", string(data)}, " ")
}

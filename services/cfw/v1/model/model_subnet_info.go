package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubnetInfo **参数解释**： 管理子网信息 **取值范围**： 不涉及
type SubnetInfo struct {

	// **参数解释**： 子网所在的可用区标识，从终端节点获取，参考[终端节点](cfw_02_0000.xml) **取值范围**： 不涉及
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// **参数解释**： 虚拟私有云下可用子网的范围 **取值范围**： 10.0.0.0/8~24 172.16.0.0/12~24 192.168.0.0/16~24
	Cidr *string `json:"cidr,omitempty"`

	// **参数解释**： 子网名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 子网id **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 子网的网关 **取值范围**： 子网网段cidr中的IP地址
	GatewayIp *string `json:"gateway_ip,omitempty"`

	// **参数解释**： 创建VPC产生的uuid **取值范围**： 不涉及
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**： 是否支持ipv6 **取值范围**： - true：是 - false：否
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`
}

func (o SubnetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubnetInfo struct{}"
	}

	return strings.Join([]string{"SubnetInfo", string(data)}, " ")
}

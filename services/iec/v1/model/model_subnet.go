package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 子网对象。
type Subnet struct {
	// 子网的ID。

	Id *string `json:"id,omitempty"`
	// 子网名称  取值范围：1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）

	Name *string `json:"name,omitempty"`
	// 子网的网段  取值范围：必须在vpc对应cidr范围内  约束：必须是cidr格式。掩码长度不能大于28

	Cidr *string `json:"cidr,omitempty"`
	// 子网dns服务器地址列表

	DnsList *[]string `json:"dnsList,omitempty"`
	// 子网的网关  取值范围：子网网段中的IP地址  约束：必须是ip格式

	GatewayIp *string `json:"gateway_ip,omitempty"`
	// 子网是否开启dhcp功能

	DhcpEnable *bool `json:"dhcp_enable,omitempty"`
	// 子网dns服务器地址1

	PrimaryDns *string `json:"primary_dns,omitempty"`
	// 子网dns服务器地址2

	SecondaryDns *string `json:"secondary_dns,omitempty"`
	// 子网的状态  取值范围： - ACTIVE：表示子网已挂载到ROUTER上 - UNKNOWN：表示子网还未挂载到ROUTER上 - ERROR：表示子网状态故障

	Status *SubnetStatus `json:"status,omitempty"`
	// 虚拟私有云ID。

	VpcId *string `json:"vpc_id,omitempty"`
	// 子网所属的站点ID。

	SiteId *string `json:"site_id,omitempty"`
	// 子网所属的站点信息。

	SiteInfo *string `json:"site_info,omitempty"`
	// 对应网络（OpenStack Neutron接口） id。

	NeutronNetworkId *string `json:"neutron_network_id,omitempty"`
	// 对应子网（OpenStack Neutron接口） id。

	NeutronSubnetId *string `json:"neutron_subnet_id,omitempty"`
}

func (o Subnet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Subnet struct{}"
	}

	return strings.Join([]string{"Subnet", string(data)}, " ")
}

type SubnetStatus struct {
	value string
}

type SubnetStatusEnum struct {
	ACTIVE  SubnetStatus
	UNKNOWN SubnetStatus
	ERROR   SubnetStatus
}

func GetSubnetStatusEnum() SubnetStatusEnum {
	return SubnetStatusEnum{
		ACTIVE: SubnetStatus{
			value: "ACTIVE",
		},
		UNKNOWN: SubnetStatus{
			value: "UNKNOWN",
		},
		ERROR: SubnetStatus{
			value: "ERROR",
		},
	}
}

func (c SubnetStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubnetStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PrePaidServerNic 待创建云服务器的网卡信息。
type PrePaidServerNic struct {

	// 待创建云服务器所在的子网信息，需要指定vpcid对应VPC下的子网ID，UUID格式。  可以通过VPC服务 [查询子网](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=VPC&api=ListSubnets) 接口查询，该接口支持通过创建云服务器填写的vpcid进行过滤查询。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 网卡ID，UUID格式。 当该字段不为空时，表示挂载指定的网卡。port_id和subnet_id不能同时为空。 网卡ID可以从虚拟私有云的“查询端口列表”章节查询到。 约束： 网卡状态必须为DOWN。 网卡的vpcid必须和传入的vpcid一致。 当port_id和subnet_id同时存在的时候，优先使用port_id。当选择port_id不为空时，代表此时使用的是弹性网卡，此时security_groups和ip_address等参数不生效。
	PortId *string `json:"port_id,omitempty"`

	// 待创建云服务器网卡的IP地址，IPv4格式。  约束：  - 不填或空字符串，默认在子网（subnet）中自动分配一个未使用的IP作网卡的IP地址。 - 若指定IP地址，该IP地址必须在子网（subnet）对应的网段内，且未被使用。
	IpAddress *string `json:"ip_address,omitempty"`

	// 是否支持ipv6。  取值为true时，标识此网卡支持ipv6。
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	Ipv6Bandwidth *PrePaidServerIpv6Bandwidth `json:"ipv6_bandwidth,omitempty"`

	// IP/Mac对列表， 约束：IP地址不允许为 “0.0.0.0/0” 如果allowed_address_pairs配置地址池较大的CIDR（掩码小于24位），建议为该port配置一个单独的安全组 如果allowed_address_pairs为“1.1.1.1/0”，表示关闭源目地址检查开关 被绑定的云服务器网卡allowed_address_pairs填“1.1.1.1/0”
	AllowedAddressPairs *[]CreateServerNicAllowedAddressPairs `json:"allowed_address_pairs,omitempty"`
}

func (o PrePaidServerNic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrePaidServerNic struct{}"
	}

	return strings.Join([]string{"PrePaidServerNic", string(data)}, " ")
}

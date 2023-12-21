package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubnetOption 创建子网对象数据。
type CreateSubnetOption struct {

	// 子网名称  约束：由中文字符、字母、数字、中划线和下划线和点组成，长度为1~64个字符
	Name string `json:"name"`

	// 子网的网段  取值范围：必须在vpc对应cidr范围内  约束：必须是cidr格式。掩码长度不能大于28
	Cidr string `json:"cidr"`

	// 子网的网关  取值范围：子网网段中的IP地址  约束：必须是ip格式
	GatewayIp string `json:"gateway_ip"`

	// 子网是否开启dhcp功能  取值范围：true（开启），false（关闭）  约束：不填时默认为true。当设置为false时，会导致新创建的ECS无法获取IP地址，cloudinit无法注入账号密码，请谨慎操作。
	DhcpEnable *bool `json:"dhcp_enable,omitempty"`

	// 子网dns服务器地址1  约束：ip格式，不支持IPv6地址
	PrimaryDns *string `json:"primary_dns,omitempty"`

	// 子网dns服务器地址2  约束：ip格式，不支持IPv6地址
	SecondaryDns *string `json:"secondary_dns,omitempty"`

	// 子网dns服务器地址的集合；如果想使用两个以上dns服务器，请使用该字段  约束：是子网dns服务器地址1跟子网dns服务器地址2的合集的父集，不支持IPv6地址。
	DnsList *[]string `json:"dnsList,omitempty"`

	// 子网所在VPC的ID。
	VpcId string `json:"vpc_id"`

	// 子网归属的站点ID,从站点信息列表中获取。
	SiteId string `json:"site_id"`
}

func (o CreateSubnetOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubnetOption struct{}"
	}

	return strings.Join([]string{"CreateSubnetOption", string(data)}, " ")
}

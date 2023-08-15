package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubnetOption 更新子网的结构体。
type UpdateSubnetOption struct {

	// 子网名称  取值范围：0-64，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`

	// 子网是否开启dhcp功能  取值范围：true（开启），false（关闭）  约束：不填时默认为true。当设置为false时，会导致新创建的实例无法获取IP地址，cloudinit无法注入帐号密码，请谨慎操作。
	DhcpEnable *bool `json:"dhcp_enable,omitempty"`

	// 子网dns服务器地址1  约束：ip格式
	PrimaryDns *string `json:"primary_dns,omitempty"`

	// 子网dns服务器地址2  约束：ip格式
	SecondaryDns *string `json:"secondary_dns,omitempty"`

	// 子网dns服务器地址的集合；如果想使用两个以上dns服务器，请使用该字段。  约束：是子网dns服务器地址1跟子网dns服务器地址2的合集的父集
	DnsList *[]string `json:"dnsList,omitempty"`
}

func (o UpdateSubnetOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubnetOption struct{}"
	}

	return strings.Join([]string{"UpdateSubnetOption", string(data)}, " ")
}

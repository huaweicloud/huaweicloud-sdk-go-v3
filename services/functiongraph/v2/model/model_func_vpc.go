package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FuncVpc 函数vpc配置。需同时为函数配置具有vpc权限的委托。
type FuncVpc struct {

	// 域名id。
	DomainId *string `json:"domain_id,omitempty"`

	// 租户的project id。
	Namespace *string `json:"namespace,omitempty"`

	// 虚拟私有云名称。
	VpcName *string `json:"vpc_name,omitempty"`

	// 虚拟私有云唯一标识。
	VpcId string `json:"vpc_id"`

	// 子网名称。
	SubnetName *string `json:"subnet_name,omitempty"`

	// 子网编号。
	SubnetId string `json:"subnet_id"`

	// 子网掩码。
	Cidr *string `json:"cidr,omitempty"`

	// 网关。
	Gateway *string `json:"gateway,omitempty"`

	// 安全组
	SecurityGroups *[]string `json:"security_groups,omitempty"`
}

func (o FuncVpc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FuncVpc struct{}"
	}

	return strings.Join([]string{"FuncVpc", string(data)}, " ")
}

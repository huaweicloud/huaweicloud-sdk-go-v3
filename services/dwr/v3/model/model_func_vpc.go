package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FuncVpc 函数vpc
type FuncVpc struct {

	// 子网编号。当func_vpc非空时必选。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 虚拟私有云唯一标识。当func_vpc非空时必选。
	VpcId *string `json:"vpc_id,omitempty"`

	// VPC名称。
	VpcName *string `json:"vpc_name,omitempty"`

	// 子网名称。
	SubnetName *string `json:"subnet_name,omitempty"`

	// 子网掩码。
	Cidr *string `json:"cidr,omitempty"`

	// 网关。
	Gateway *string `json:"gateway,omitempty"`
}

func (o FuncVpc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FuncVpc struct{}"
	}

	return strings.Join([]string{"FuncVpc", string(data)}, " ")
}

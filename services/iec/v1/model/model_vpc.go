package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Vpc 虚拟私有云的数据对象。
type Vpc struct {

	// 虚拟私有云的ID。
	Id *string `json:"id,omitempty"`

	// 虚拟私有云名称  取值范围：1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）  约束：同一个帐号下的名称不能重复
	Name *string `json:"name,omitempty"`

	// 虚拟私有云下可用子网的范围  取值范围： 10.0.0.0/8~24 172.16.0.0/12~24 192.168.0.0/16~24  约束：必须是cidr格式，例如:192.168.0.0/16
	Cidr *string `json:"cidr,omitempty"`

	// 虚拟私有云的模式。
	Mode *string `json:"mode,omitempty"`

	// 子网的数目。
	SubnetNum *int32 `json:"subnet_num,omitempty"`
}

func (o Vpc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Vpc struct{}"
	}

	return strings.Join([]string{"Vpc", string(data)}, " ")
}

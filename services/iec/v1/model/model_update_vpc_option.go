package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpcOption 更新Vpc参数
type UpdateVpcOption struct {

	// 虚拟私有云名称  取值范围：0-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）  约束：同一个帐号下的VPC不允许重名
	Name *string `json:"name,omitempty"`

	// 虚拟私有云下可用子网的范围。  约束条件：  SYSTEM模式，cidr取值范围：10.0.0.0/8~10.255.0.0/16或者172.16.0.0/12 ~ 172.31.0.0/16或者192.168.0.0/16 。 [CUSTOMER模式，cidr的取值范围：10.0.0.0/8~10.255.255.0/24或者172.16.0.0/12 ~ 172.32.255.0/24或者192.168.0.0/16~192.168.255.0/24。](tag:internal)
	Cidr *string `json:"cidr,omitempty"`
}

func (o UpdateVpcOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcOption struct{}"
	}

	return strings.Join([]string{"UpdateVpcOption", string(data)}, " ")
}

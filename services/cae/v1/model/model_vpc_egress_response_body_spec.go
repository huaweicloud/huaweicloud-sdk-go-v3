package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VpcEgressResponseBodySpec 创建CAE环境访问VPC配置。
type VpcEgressResponseBodySpec struct {

	// CAE环境VPCID。
	VpcId *string `json:"vpc_id,omitempty"`

	// CAE环境子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// CAE环境访问VPC配置。
	Cidrs *[]EgressCidr `json:"cidrs,omitempty"`
}

func (o VpcEgressResponseBodySpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcEgressResponseBodySpec struct{}"
	}

	return strings.Join([]string{"VpcEgressResponseBodySpec", string(data)}, " ")
}

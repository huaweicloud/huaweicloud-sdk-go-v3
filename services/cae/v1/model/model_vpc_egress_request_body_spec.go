package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VpcEgressRequestBodySpec 创建CAE环境访问VPC配置。
type VpcEgressRequestBodySpec struct {

	// CAE环境访问VPC配置。
	Cidrs []EgressCidr `json:"cidrs"`
}

func (o VpcEgressRequestBodySpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcEgressRequestBodySpec struct{}"
	}

	return strings.Join([]string{"VpcEgressRequestBodySpec", string(data)}, " ")
}

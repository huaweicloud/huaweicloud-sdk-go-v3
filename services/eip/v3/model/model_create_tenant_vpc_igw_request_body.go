package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTenantVpcIgwRequestBody 创建虚拟IGW的请求体
type CreateTenantVpcIgwRequestBody struct {
	VpcIgw *CreateTenantVpcIgwRequestBodyVpcIgw `json:"vpc_igw"`
}

func (o CreateTenantVpcIgwRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTenantVpcIgwRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTenantVpcIgwRequestBody", string(data)}, " ")
}

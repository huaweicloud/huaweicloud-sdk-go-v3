package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTenantVpcIgwRequestBody 修改虚拟IGW的请求体
type UpdateTenantVpcIgwRequestBody struct {
	VpcIgw *UpdateTenantVpcIgwRequestBodyVpcIgw `json:"vpc_igw"`
}

func (o UpdateTenantVpcIgwRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTenantVpcIgwRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTenantVpcIgwRequestBody", string(data)}, " ")
}

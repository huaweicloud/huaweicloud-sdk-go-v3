package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTenantVpcIgwRequest Request Object
type DeleteTenantVpcIgwRequest struct {

	// 虚拟igw的uuid
	VpcIgwId string `json:"vpc_igw_id"`
}

func (o DeleteTenantVpcIgwRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTenantVpcIgwRequest struct{}"
	}

	return strings.Join([]string{"DeleteTenantVpcIgwRequest", string(data)}, " ")
}

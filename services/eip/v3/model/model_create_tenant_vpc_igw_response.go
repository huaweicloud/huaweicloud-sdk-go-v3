package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTenantVpcIgwResponse Response Object
type CreateTenantVpcIgwResponse struct {
	VpcIgw *VpcIgwsTenantResp `json:"vpc_igw,omitempty"`

	// 本次请求编号
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTenantVpcIgwResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTenantVpcIgwResponse struct{}"
	}

	return strings.Join([]string{"CreateTenantVpcIgwResponse", string(data)}, " ")
}

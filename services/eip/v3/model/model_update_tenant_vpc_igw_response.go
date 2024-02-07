package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTenantVpcIgwResponse Response Object
type UpdateTenantVpcIgwResponse struct {
	VpcIgw *VpcIgwsTenantResp `json:"vpc_igw,omitempty"`

	// 本次请求编号
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTenantVpcIgwResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTenantVpcIgwResponse struct{}"
	}

	return strings.Join([]string{"UpdateTenantVpcIgwResponse", string(data)}, " ")
}

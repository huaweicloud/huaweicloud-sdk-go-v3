package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInternalVpcIgwResponse Response Object
type ShowInternalVpcIgwResponse struct {
	VpcIgw *VpcIgwsTenantResp `json:"vpc_igw,omitempty"`

	// 本次请求编号
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInternalVpcIgwResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInternalVpcIgwResponse struct{}"
	}

	return strings.Join([]string{"ShowInternalVpcIgwResponse", string(data)}, " ")
}

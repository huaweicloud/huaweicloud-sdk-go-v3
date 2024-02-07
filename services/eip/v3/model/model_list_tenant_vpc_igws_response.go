package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantVpcIgwsResponse Response Object
type ListTenantVpcIgwsResponse struct {

	// 虚拟IGW列表对象
	VpcIgws *[]VpcIgwsTenantResp `json:"vpc_igws,omitempty"`

	// 本次请求编号
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTenantVpcIgwsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantVpcIgwsResponse struct{}"
	}

	return strings.Join([]string{"ListTenantVpcIgwsResponse", string(data)}, " ")
}

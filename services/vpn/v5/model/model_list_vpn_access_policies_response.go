package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpnAccessPoliciesResponse Response Object
type ListVpnAccessPoliciesResponse struct {

	// 访问资源策略信息
	AccessPolicies *[]VpnAccessPolicy `json:"access_policies,omitempty"`

	// 总数
	TotalCount *int32 `json:"total_count,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListVpnAccessPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpnAccessPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListVpnAccessPoliciesResponse", string(data)}, " ")
}

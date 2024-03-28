package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrganizationPoliciesResponse Response Object
type ListOrganizationPoliciesResponse struct {

	// 组织策略列表
	Policies *[]OrganizationPolicy `json:"policies,omitempty"`

	// 组织策略数量
	Count *int32 `json:"count,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询
	Offset         *int32 `json:"offset,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListOrganizationPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrganizationPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListOrganizationPoliciesResponse", string(data)}, " ")
}

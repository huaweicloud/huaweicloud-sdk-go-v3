package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrganizationPolicyDetailResponse Response Object
type ListOrganizationPolicyDetailResponse struct {

	// 组织策略部署状态列表
	Policies *[]OrganizationPolicyStatus `json:"policies,omitempty"`

	// 组织策略状态成员数量
	Count *int32 `json:"count,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询
	Offset         *int32 `json:"offset,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListOrganizationPolicyDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrganizationPolicyDetailResponse struct{}"
	}

	return strings.Join([]string{"ListOrganizationPolicyDetailResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityDynamicMaskingPoliciesResponse Response Object
type ListSecurityDynamicMaskingPoliciesResponse struct {

	// 动态脱敏策略总条数。
	Total *int32 `json:"total,omitempty"`

	// 动态数据脱敏策略列表。
	Policies       *[]DynamicMaskingPolicySet `json:"policies,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListSecurityDynamicMaskingPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityDynamicMaskingPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityDynamicMaskingPoliciesResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceSignPoliciesResponse Response Object
type ListInstanceSignPoliciesResponse struct {

	// 签名策略列表
	Policies *[]SignPolicyDetail `json:"policies,omitempty"`

	// 签名策略总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceSignPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceSignPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceSignPoliciesResponse", string(data)}, " ")
}

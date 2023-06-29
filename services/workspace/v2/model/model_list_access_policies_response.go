package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessPoliciesResponse Response Object
type ListAccessPoliciesResponse struct {

	// 查询接入策略响应。
	Policies *[]AccessPolicyDetailInfo `json:"policies,omitempty"`

	// 策略总数。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAccessPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListAccessPoliciesResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIncentiveDiscountPoliciesResponse Response Object
type ListIncentiveDiscountPoliciesResponse struct {

	// 产品折扣和激励策略信息列表。 具体请参见表2。
	Policies *[]IncentiveAndDiscountPolicy `json:"policies,omitempty"`

	// 查询总条数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListIncentiveDiscountPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIncentiveDiscountPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListIncentiveDiscountPoliciesResponse", string(data)}, " ")
}

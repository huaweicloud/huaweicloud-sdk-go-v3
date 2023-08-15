package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlavorSellPoliciesResponse Response Object
type ListFlavorSellPoliciesResponse struct {

	// 云服务器规格销售策略数量。
	Count *int32 `json:"count,omitempty"`

	// 云服务器规格销售策略。
	SellPolicies   *[]ListFlavorSellPoliciesResult `json:"sell_policies,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListFlavorSellPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorSellPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListFlavorSellPoliciesResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCustomersBalancesDetailResponse struct {
	// 账户余额列表。只有成功的时候才返回。 具体请参见表2。

	CustomerBalances *[]CustomerBalancesV2 `json:"customer_balances,omitempty"`
	HttpStatusCode   int                   `json:"-"`
}

func (o ListCustomersBalancesDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomersBalancesDetailResponse struct{}"
	}

	return strings.Join([]string{"ListCustomersBalancesDetailResponse", string(data)}, " ")
}

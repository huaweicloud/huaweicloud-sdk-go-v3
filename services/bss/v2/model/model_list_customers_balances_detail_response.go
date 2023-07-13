package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomersBalancesDetailResponse Response Object
type ListCustomersBalancesDetailResponse struct {

	// 账户余额列表。只有成功的时候才返回。 此列表不包含非代售类子客户的数据。批量查询客户余额时，若入参携带了非代售类子客户，会被过滤。 具体请参见表2。
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

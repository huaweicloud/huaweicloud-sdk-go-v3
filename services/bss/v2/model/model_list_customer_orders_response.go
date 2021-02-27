package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListCustomerOrdersResponse struct {
	// 大于等于0的整数 符合条件的记录总数。
	TotalCount *int32 `json:"total_count,omitempty"`
	// - 客户订单详情信息。 具体请参见表2
	OrderInfos     *[]CustomerOrderV2 `json:"order_infos,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListCustomerOrdersResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCustomerOrdersResponse struct{}"
	}

	return strings.Join([]string{"ListCustomerOrdersResponse", string(data)}, " ")
}

package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListOrderDiscountsResponse struct {
	// 可用的折扣列表。 具体请参见表2。

	Discounts      *[]DiscountInfoV3 `json:"discounts,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListOrderDiscountsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrderDiscountsResponse struct{}"
	}

	return strings.Join([]string{"ListOrderDiscountsResponse", string(data)}, " ")
}

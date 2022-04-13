package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListProPricePlansResponse struct {
	// 每页的记录数

	Limit *int64 `json:"limit,omitempty"`
	// 页码，最小值是1，最大值为1000000。默认值是1.

	Offset *int64 `json:"offset,omitempty"`
	// 记录总数

	Count *int64 `json:"count,omitempty"`
	// 套餐列表

	PricePlans     *[]ProPricePlanVo `json:"price_plans,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListProPricePlansResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProPricePlansResponse struct{}"
	}

	return strings.Join([]string{"ListProPricePlansResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataArtsStudioInstancesResponse Response Object
type ListDataArtsStudioInstancesResponse struct {

	// 是否需要账单
	BillingCheck *bool `json:"billing_check,omitempty"`

	// 返回记录总数
	Count *int32 `json:"count,omitempty"`

	// 返回实例列表
	CommodityOrderLists *[]ApigCommodityOrder `json:"commodity_order_lists,omitempty"`
	HttpStatusCode      int                   `json:"-"`
}

func (o ListDataArtsStudioInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataArtsStudioInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListDataArtsStudioInstancesResponse", string(data)}, " ")
}

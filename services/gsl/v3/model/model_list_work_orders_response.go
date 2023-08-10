package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkOrdersResponse Response Object
type ListWorkOrdersResponse struct {

	// 每页的记录数
	Limit *int64 `json:"limit,omitempty"`

	// 页码，最小值是1，最大值为1000000。默认值是1.
	Offset *int64 `json:"offset,omitempty"`

	// 记录总数
	Count *int64 `json:"count,omitempty"`

	// 业务受理单列表
	WorkOrders     *[]WorkOrderVo `json:"work_orders,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListWorkOrdersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkOrdersResponse struct{}"
	}

	return strings.Join([]string{"ListWorkOrdersResponse", string(data)}, " ")
}

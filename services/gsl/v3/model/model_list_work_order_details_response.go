package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkOrderDetailsResponse Response Object
type ListWorkOrderDetailsResponse struct {

	// 每页的记录数
	Limit *int64 `json:"limit,omitempty"`

	// 页码，最小值是1，最大值为1000000。默认值是1.
	Offset *int64 `json:"offset,omitempty"`

	// 记录总数
	Count *int64 `json:"count,omitempty"`

	// 业务受理明细列表
	WorkOrderDetails *[]WorkOrderDetailVo `json:"work_order_details,omitempty"`
	HttpStatusCode   int                  `json:"-"`
}

func (o ListWorkOrderDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkOrderDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListWorkOrderDetailsResponse", string(data)}, " ")
}

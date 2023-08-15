package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkOrderDetailsRequest Request Object
type ListWorkOrderDetailsRequest struct {

	// 业务受理ID
	WorkOrderId int64 `json:"work_order_id"`

	// 容器ID
	MainSearchKey *string `json:"main_search_key,omitempty"`

	// 每页的记录数
	Limit *int64 `json:"limit,omitempty"`

	// 页码，最小值是1，最大值为1000000。默认值是1
	Offset *int64 `json:"offset,omitempty"`

	// SIM卡类型:3.实体卡
	SimType *int32 `json:"sim_type,omitempty"`

	// 业务受理明细状态：1成功、2处理中、3失败
	Status *int32 `json:"status,omitempty"`
}

func (o ListWorkOrderDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkOrderDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListWorkOrderDetailsRequest", string(data)}, " ")
}

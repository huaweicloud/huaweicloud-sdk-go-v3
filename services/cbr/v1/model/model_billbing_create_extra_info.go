package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type BillbingCreateExtraInfo struct {

	// 组合创建ID，组合创建时必传。
	CombinedOrderId *string `json:"combined_order_id,omitempty"`

	// 组合创建数量，组合创建时必填。
	CombinedOrderEcsNum *int32 `json:"combined_order_ecs_num,omitempty"`
}

func (o BillbingCreateExtraInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BillbingCreateExtraInfo struct{}"
	}

	return strings.Join([]string{"BillbingCreateExtraInfo", string(data)}, " ")
}

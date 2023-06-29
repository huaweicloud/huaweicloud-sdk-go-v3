package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CombinedOrder 组合订单
type CombinedOrder struct {

	// 组合订单 ID
	CombinedOrderId *string `json:"combined_order_id,omitempty"`

	// 组合订单中 ECS 服务器数量，当前批量最大为 500。  最小值：1  最大值：1000
	CombinedOrderEcsNum *int32 `json:"combined_order_ecs_num,omitempty"`

	// 组合订单数量。  最小值：1  最大值：1000
	CombinedOrderNum *int32 `json:"combined_order_num,omitempty"`
}

func (o CombinedOrder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CombinedOrder struct{}"
	}

	return strings.Join([]string{"CombinedOrder", string(data)}, " ")
}

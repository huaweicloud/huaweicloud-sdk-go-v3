package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ResetSimCardResponse struct {
	// 业务受理单号

	WorkOrderId *int64 `json:"work_order_id,omitempty"`
	// 套餐列表

	SimPricePlanList *[]SimPricePlanInfoVo `json:"sim_price_plan_list,omitempty"`
	HttpStatusCode   int                   `json:"-"`
}

func (o ResetSimCardResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResetSimCardResponse struct{}"
	}

	return strings.Join([]string{"ResetSimCardResponse", string(data)}, " ")
}

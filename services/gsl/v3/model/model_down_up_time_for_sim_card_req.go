package model

import (
	"encoding/json"

	"strings"
)

type DownUpTimeForSimCardReq struct {
	// 套餐列表

	PricePlanList *[]SimPricePlanInfoVo `json:"price_plan_list,omitempty"`
	// 启用停用开关

	DownUpSwitch *int32 `json:"down_up_switch,omitempty"`
}

func (o DownUpTimeForSimCardReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DownUpTimeForSimCardReq struct{}"
	}

	return strings.Join([]string{"DownUpTimeForSimCardReq", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OrderForm struct {

	// 订单数量
	Amount int32 `json:"amount"`

	// 订到已使用数量
	UsedAmount int32 `json:"used_amount"`

	// cbc订单Id
	CbcOrderId string `json:"cbc_order_id"`

	// 设备类别
	DeviceType string `json:"device_type"`

	// 订单Id
	Id string `json:"id"`

	// 订单更新时间
	OrderUpdateTime string `json:"order_update_time"`

	// 订单状态
	Status int32 `json:"status"`
}

func (o OrderForm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderForm struct{}"
	}

	return strings.Join([]string{"OrderForm", string(data)}, " ")
}

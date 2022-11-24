package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 计费类型信息，支持包年包月和按需，默认为按需。
type MysqlInstanceChargeInfo struct {

	// 计费模式。
	ChargeMode string `json:"charge_mode"`

	// 订单号。
	OrderId *string `json:"order_id,omitempty"`
}

func (o MysqlInstanceChargeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlInstanceChargeInfo struct{}"
	}

	return strings.Join([]string{"MysqlInstanceChargeInfo", string(data)}, " ")
}

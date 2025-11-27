package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNatGatewayOptionPrepaidOptions 支付属性
type UpdateNatGatewayOptionPrepaidOptions struct {

	// 是否自动支付
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o UpdateNatGatewayOptionPrepaidOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewayOptionPrepaidOptions struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewayOptionPrepaidOptions", string(data)}, " ")
}

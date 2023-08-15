package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstancesPeriodOrderResponse Response Object
type CreateInstancesPeriodOrderResponse struct {

	// 描述
	Description *string `json:"description,omitempty"`

	// 返回码
	Code *string `json:"code,omitempty"`

	// 订单ID
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstancesPeriodOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstancesPeriodOrderResponse struct{}"
	}

	return strings.Join([]string{"CreateInstancesPeriodOrderResponse", string(data)}, " ")
}

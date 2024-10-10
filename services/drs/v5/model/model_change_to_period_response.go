package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeToPeriodResponse Response Object
type ChangeToPeriodResponse struct {

	// 订单ID。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeToPeriodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeToPeriodResponse struct{}"
	}

	return strings.Join([]string{"ChangeToPeriodResponse", string(data)}, " ")
}

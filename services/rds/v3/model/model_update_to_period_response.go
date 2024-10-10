package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateToPeriodResponse Response Object
type UpdateToPeriodResponse struct {

	// 转包周期订单ID
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateToPeriodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateToPeriodResponse struct{}"
	}

	return strings.Join([]string{"UpdateToPeriodResponse", string(data)}, " ")
}

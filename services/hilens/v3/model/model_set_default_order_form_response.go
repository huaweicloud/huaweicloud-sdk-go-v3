package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetDefaultOrderFormResponse Response Object
type SetDefaultOrderFormResponse struct {

	// 订单ID
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetDefaultOrderFormResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetDefaultOrderFormResponse struct{}"
	}

	return strings.Join([]string{"SetDefaultOrderFormResponse", string(data)}, " ")
}

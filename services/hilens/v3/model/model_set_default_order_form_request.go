package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetDefaultOrderFormRequest Request Object
type SetDefaultOrderFormRequest struct {

	// 订单ID
	OrderId string `json:"order_id"`
}

func (o SetDefaultOrderFormRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetDefaultOrderFormRequest struct{}"
	}

	return strings.Join([]string{"SetDefaultOrderFormRequest", string(data)}, " ")
}

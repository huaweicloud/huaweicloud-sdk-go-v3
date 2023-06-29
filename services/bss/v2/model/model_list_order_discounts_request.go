package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrderDiscountsRequest Request Object
type ListOrderDiscountsRequest struct {

	// 订单ID。
	OrderId string `json:"order_id"`
}

func (o ListOrderDiscountsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrderDiscountsRequest struct{}"
	}

	return strings.Join([]string{"ListOrderDiscountsRequest", string(data)}, " ")
}

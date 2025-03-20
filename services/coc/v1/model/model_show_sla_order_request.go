package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlaOrderRequest Request Object
type ShowSlaOrderRequest struct {

	// 工单ID
	OrderId string `json:"order_id"`
}

func (o ShowSlaOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlaOrderRequest struct{}"
	}

	return strings.Join([]string{"ShowSlaOrderRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesktopBatchOrderResponse Response Object
type CreateDesktopBatchOrderResponse struct {

	// 批量生成订单结果。
	Orders         *[]OrderV5 `json:"orders,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o CreateDesktopBatchOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopBatchOrderResponse struct{}"
	}

	return strings.Join([]string{"CreateDesktopBatchOrderResponse", string(data)}, " ")
}

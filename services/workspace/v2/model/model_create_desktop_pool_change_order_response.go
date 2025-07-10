package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesktopPoolChangeOrderResponse Response Object
type CreateDesktopPoolChangeOrderResponse struct {

	// 批量生成订单结果。
	Orders         *[]OrderV5 `json:"orders,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o CreateDesktopPoolChangeOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopPoolChangeOrderResponse struct{}"
	}

	return strings.Join([]string{"CreateDesktopPoolChangeOrderResponse", string(data)}, " ")
}

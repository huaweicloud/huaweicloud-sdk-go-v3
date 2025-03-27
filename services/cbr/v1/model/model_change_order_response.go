package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeOrderResponse Response Object
type ChangeOrderResponse struct {

	// 订单ID
	OrderId *string `json:"orderId,omitempty"`

	// 变更状态码
	RetCode float32 `json:"retCode,omitempty"`

	// 变更信息
	RetMsg         *string `json:"retMsg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeOrderResponse struct{}"
	}

	return strings.Join([]string{"ChangeOrderResponse", string(data)}, " ")
}

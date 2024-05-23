package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeInstanceTypeResponse Response Object
type ChangeInstanceTypeResponse struct {

	// 订单ID。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeInstanceTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeInstanceTypeResponse struct{}"
	}

	return strings.Join([]string{"ChangeInstanceTypeResponse", string(data)}, " ")
}

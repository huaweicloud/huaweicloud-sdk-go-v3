package model

import (
	"encoding/json"

	"strings"
)

type MysqlChangeSpecificationRequest struct {
	ResizeFlavor *MysqlResizeFlavor `json:"resize_flavor"`
	// 表示是否自动从客户的账户中支付。true，表示自动支付，默认该方式。false，表示手动支付。

	IsAutoPay *string `json:"is_auto_pay,omitempty"`
}

func (o MysqlChangeSpecificationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MysqlChangeSpecificationRequest struct{}"
	}

	return strings.Join([]string{"MysqlChangeSpecificationRequest", string(data)}, " ")
}

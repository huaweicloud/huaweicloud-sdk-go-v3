package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeServerChargeModeResponse Response Object
type ChangeServerChargeModeResponse struct {

	// 按需转包提交后返回的订单ID，用户可以使用该ID对订单结果进行查询。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeServerChargeModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerChargeModeResponse struct{}"
	}

	return strings.Join([]string{"ChangeServerChargeModeResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeVaultChargeModeResponse Response Object
type ChangeVaultChargeModeResponse struct {

	// 订单ID
	OrderId *string `json:"orderId,omitempty"`

	// 变更状态码
	RetCode float32 `json:"retCode,omitempty"`

	// 变更信息
	RetMsg         *string `json:"retMsg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeVaultChargeModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeVaultChargeModeResponse struct{}"
	}

	return strings.Join([]string{"ChangeVaultChargeModeResponse", string(data)}, " ")
}

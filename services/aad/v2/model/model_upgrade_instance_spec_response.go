package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeInstanceSpecResponse Response Object
type UpgradeInstanceSpecResponse struct {

	// 订单id
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpgradeInstanceSpecResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeInstanceSpecResponse struct{}"
	}

	return strings.Join([]string{"UpgradeInstanceSpecResponse", string(data)}, " ")
}

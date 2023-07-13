package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangePrepaidCloudWafResponse Response Object
type ChangePrepaidCloudWafResponse struct {

	// 订单id
	OrderId        *string `json:"orderId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangePrepaidCloudWafResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePrepaidCloudWafResponse struct{}"
	}

	return strings.Join([]string{"ChangePrepaidCloudWafResponse", string(data)}, " ")
}

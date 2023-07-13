package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrepaidCloudWafResponse Response Object
type CreatePrepaidCloudWafResponse struct {

	// 订单id
	OrderId        *string `json:"orderId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePrepaidCloudWafResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrepaidCloudWafResponse struct{}"
	}

	return strings.Join([]string{"CreatePrepaidCloudWafResponse", string(data)}, " ")
}

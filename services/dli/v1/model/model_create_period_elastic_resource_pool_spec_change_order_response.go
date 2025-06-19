package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePeriodElasticResourcePoolSpecChangeOrderResponse Response Object
type CreatePeriodElasticResourcePoolSpecChangeOrderResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	// 订单id
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePeriodElasticResourcePoolSpecChangeOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePeriodElasticResourcePoolSpecChangeOrderResponse struct{}"
	}

	return strings.Join([]string{"CreatePeriodElasticResourcePoolSpecChangeOrderResponse", string(data)}, " ")
}

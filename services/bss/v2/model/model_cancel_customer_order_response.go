package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelCustomerOrderResponse Response Object
type CancelCustomerOrderResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelCustomerOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelCustomerOrderResponse struct{}"
	}

	return strings.Join([]string{"CancelCustomerOrderResponse", string(data)}, " ")
}

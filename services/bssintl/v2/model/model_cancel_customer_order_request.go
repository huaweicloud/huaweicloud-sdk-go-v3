package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelCustomerOrderRequest Request Object
type CancelCustomerOrderRequest struct {
	Body *CancelCustomerOrderReq `json:"body,omitempty"`
}

func (o CancelCustomerOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelCustomerOrderRequest struct{}"
	}

	return strings.Join([]string{"CancelCustomerOrderRequest", string(data)}, " ")
}

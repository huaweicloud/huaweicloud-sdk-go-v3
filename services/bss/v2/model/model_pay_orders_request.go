package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type PayOrdersRequest struct {
	Body *PayCustomerOrderReq `json:"body,omitempty"`
}

func (o PayOrdersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PayOrdersRequest struct{}"
	}

	return strings.Join([]string{"PayOrdersRequest", string(data)}, " ")
}

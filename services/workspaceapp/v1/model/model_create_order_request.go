package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrderRequest Request Object
type CreateOrderRequest struct {
	Body *CreateOrderReq `json:"body,omitempty"`
}

func (o CreateOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrderRequest struct{}"
	}

	return strings.Join([]string{"CreateOrderRequest", string(data)}, " ")
}

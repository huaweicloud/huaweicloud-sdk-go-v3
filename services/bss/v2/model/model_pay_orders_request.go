package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PayOrdersRequest Request Object
type PayOrdersRequest struct {
	Body *PayCustomerOrderV3Req `json:"body,omitempty"`
}

func (o PayOrdersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PayOrdersRequest struct{}"
	}

	return strings.Join([]string{"PayOrdersRequest", string(data)}, " ")
}

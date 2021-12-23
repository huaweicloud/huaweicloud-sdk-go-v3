package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchSetSubCustomerDiscountRequest struct {
	Body *SetPartnerDiscountsReq `json:"body,omitempty"`
}

func (o BatchSetSubCustomerDiscountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetSubCustomerDiscountRequest struct{}"
	}

	return strings.Join([]string{"BatchSetSubCustomerDiscountRequest", string(data)}, " ")
}

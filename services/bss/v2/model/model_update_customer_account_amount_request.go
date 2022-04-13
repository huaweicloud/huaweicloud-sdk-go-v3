package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateCustomerAccountAmountRequest struct {
	Body *AdjustAccountReq `json:"body,omitempty"`
}

func (o UpdateCustomerAccountAmountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCustomerAccountAmountRequest struct{}"
	}

	return strings.Join([]string{"UpdateCustomerAccountAmountRequest", string(data)}, " ")
}

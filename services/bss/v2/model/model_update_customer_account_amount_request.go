package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateCustomerAccountAmountRequest struct {
	Body *AdjustAccountReq `json:"body,omitempty" xml:"body"`
}

func (o UpdateCustomerAccountAmountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCustomerAccountAmountRequest struct{}"
	}

	return strings.Join([]string{"UpdateCustomerAccountAmountRequest", string(data)}, " ")
}

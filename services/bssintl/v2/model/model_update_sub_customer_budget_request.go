package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateSubCustomerBudgetRequest struct {
	Body *ModSubCustomerBudgetReq `json:"body,omitempty" xml:"body"`
}

func (o UpdateSubCustomerBudgetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubCustomerBudgetRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubCustomerBudgetRequest", string(data)}, " ")
}
